package keeper

import (
	"bytes"
	"reflect"
	"time"

	sdkerrors "cosmossdk.io/errors"
	"github.com/cometbft/cometbft/crypto/tmhash"
	"github.com/cometbft/cometbft/light"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibctmtypes "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
)

// ExtendedKeeper is same as the original Keeper, except that
//   - it provides hooks for notifying other modules on received headers
//   - it applies different verification rules on received headers
//     (notably, intercepting headers rather than freezing clients upon errors that indicate dishonest majority)
type ExtendedKeeper struct {
	Keeper
	hooks ClientHooks
}

// NewExtendedKeeper creates a new NewExtendedKeeper instance
func NewExtendedKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, paramSpace paramtypes.Subspace, sk types.StakingKeeper, uk types.UpgradeKeeper) ExtendedKeeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	k := Keeper{
		storeKey:      key,
		cdc:           cdc,
		paramSpace:    paramSpace,
		stakingKeeper: sk,
		upgradeKeeper: uk,
	}
	return ExtendedKeeper{
		Keeper: k,
		hooks:  nil,
	}
}

// SetHooks sets the hooks for ExtendedKeeper
func (ek *ExtendedKeeper) SetHooks(ch ClientHooks) *ExtendedKeeper {
	if ek.hooks != nil {
		panic("cannot set hooks twice")
	}
	ek.hooks = ch

	return ek
}

// UpdateClient applies all verification rules on the header before executing the original UpdateClient() logic.
// There are three possible outcomes after the verification:
// 1. All verifications are passed: timestamp the header, pass the header to original UpdateClient()
// 2. Verification fails with errors that only happen upon dishonest majority: timestamp the header, don't pass the header to original UpdateClient()
// 3. Verification fails with normal errors: don't timestamp the header, don't pass the header to original UpdateClient()
func (ek ExtendedKeeper) UpdateClient(ctx sdk.Context, clientID string, header exported.ClientMessage) error {
	// check if the client type is 07-tendermint
	clientState, found := ek.GetClientState(ctx, clientID)
	if !found {
		return sdkerrors.Wrapf(types.ErrClientNotFound, "cannot update client with ID %s", clientID)
	}
	_, isTmClient := clientState.(*ibctmtypes.ClientState)

	// Our logic only applies when
	// - header is not nil (header can be nil in the original IBC-Go design, e.g., in `BeginBlocker`)
	// - the client type is 07-tendermint (IBC-Go only supports 07-tendermint client at the moment)
	if header == nil {
		return ek.Keeper.UpdateClient(ctx, clientID, header)
	}

	// TODO to support wasm client, we obviously need to do something different here.
	// whole callback AfterHeaderWithValidCommit, must be generalized to support
	// other client types, not only tendermint.
	if !isTmClient {
		return ek.Keeper.UpdateClient(ctx, clientID, header)
	}

	switch msg := header.(type) {
	case *ibctmtypes.Header:
		// TODO in ibc-go v7 this whole verification was moved to seprarate function
		// clientState.CheckForMisbehaviour. Probably we can do a lot of simplifications
		// here.
		err := ek.checkHeader(ctx, clientID, msg)

		// good header, timestamp it on the canonical chain indexer
		if err == nil {
			txHash := tmhash.Sum(ctx.TxBytes())                    // get hash of the tx that includes this header
			ek.AfterHeaderWithValidCommit(ctx, txHash, msg, false) // invoke hooks to notify ZoneConcierge to timestamp this header
		}

		// The header has a QC but {is on a fork, its timestamp is not monotonic}, timestamp it on the fork indexer, and return to avoid freezing the light client
		// These two errors only happen upon dishonest majority
		if sdkerrors.IsOf(err, types.ErrForkedHeaderWithValidCommit, types.ErrHeaderNonMonotonicTimestamp) {
			ctx.Logger().Debug("received a header that has QC but is on a fork")
			txHash := tmhash.Sum(ctx.TxBytes())
			ek.AfterHeaderWithValidCommit(ctx, txHash, msg, true)
			return err
		}

		// Upon other errors (that happen under honest majority), return error without timestamping it
		if err != nil {
			return err
		}

		// the header has a valid QC and is extending canonical chain
		// follow the original verification rules to update client state
		return ek.Keeper.UpdateClient(ctx, clientID, header)
	case *ibctmtypes.Misbehaviour:
		// on Misbehaviour,just forward to standard client logic
		return ek.Keeper.UpdateClient(ctx, clientID, header)
	default:
		return types.ErrInvalidClientType
	}
}

// checkHeader checks
// - if a header is on fork,
// - if a header has a valid QC or not, and
// - other miscellaneous verification rules
// It is essentially `CheckHeaderAndUpdateState` without updating the state
// (adapted from https://github.com/cosmos/ibc-go/blob/v5.0.0/modules/core/02-client/keeper/client.go#L60)
func (ek ExtendedKeeper) checkHeader(ctx sdk.Context, clientID string, tmHeader *ibctmtypes.Header) error {
	clientState, found := ek.GetClientState(ctx, clientID)
	if !found {
		return sdkerrors.Wrapf(types.ErrClientNotFound, "cannot update client with ID %s", clientID)
	}

	clientStore := ek.ClientStore(ctx, clientID)
	if status := clientState.Status(ctx, clientStore, ek.cdc); status != exported.Active {
		return sdkerrors.Wrapf(types.ErrClientNotActive, "cannot update client (%s) with status %s", clientID, status)
	}

	// Check if the header is on a fork
	// If the consensus state exists, and it does not match this header, then this header is on a fork
	// (adapted from https://github.com/cosmos/ibc-go/blob/v5.0.0/modules/light-clients/07-tendermint/types/update.go#L63-L77)
	isOnFork := false
	prevConsState, _ := ibctmtypes.GetConsensusState(clientStore, ek.cdc, tmHeader.GetHeight())
	if prevConsState != nil {
		if !reflect.DeepEqual(prevConsState, tmHeader.ConsensusState()) {
			isOnFork = true
		}
	}

	// Check if the header can pass all verification rules, including the QC
	// Note that the first header on a fork can also be valid
	// (adapted from https://github.com/cosmos/ibc-go/blob/v5.0.0/modules/light-clients/07-tendermint/types/update.go#L79-L89)
	trustedConsState, exists := ibctmtypes.GetConsensusState(clientStore, ek.cdc, tmHeader.TrustedHeight)
	if !exists {
		return sdkerrors.Wrapf(
			types.ErrConsensusStateNotFound, "could not get consensus state from clientstore at TrustedHeight: %s", tmHeader.TrustedHeight,
		)
	}
	// asserting clientState to that of Tendermint client
	cs, ok := clientState.(*ibctmtypes.ClientState)
	if !ok {
		return sdkerrors.Wrapf(
			types.ErrFailedClientStateVerification, "expected type %T, got %T", &ibctmtypes.ClientState{}, clientState,
		)
	}
	if err := checkValidity(cs, trustedConsState, tmHeader, ctx.BlockTime()); err != nil {
		return err
	}

	// this header passes QC verifications but is on a fork
	if isOnFork {
		return types.ErrForkedHeaderWithValidCommit
	}

	consState := tmHeader.ConsensusState()
	// Check that consensus state timestamps are monotonic
	prevCons, prevOk := ibctmtypes.GetPreviousConsensusState(clientStore, ek.cdc, tmHeader.GetHeight())
	nextCons, nextOk := ibctmtypes.GetNextConsensusState(clientStore, ek.cdc, tmHeader.GetHeight())
	// if previous consensus state exists, check consensus state time is greater than previous consensus state time
	// if previous consensus state is not before current consensus state, freeze the client and return.
	if prevOk && !prevCons.Timestamp.Before(consState.Timestamp) {
		return types.ErrHeaderNonMonotonicTimestamp
	}
	// if next consensus state exists, check consensus state time is less than next consensus state time
	// if next consensus state is not after current consensus state, freeze the client and return.
	if nextOk && !nextCons.Timestamp.After(consState.Timestamp) {
		return types.ErrHeaderNonMonotonicTimestamp
	}

	return nil
}

// checkValidity checks if the Tendermint header is valid.
// CONTRACT: consState.Height == header.TrustedHeight
// (copied from https://github.com/cosmos/ibc-go/blob/v5.0.0/modules/light-clients/07-tendermint/types/update.go#L168-L248)
func checkValidity(
	clientState *ibctmtypes.ClientState, consState *ibctmtypes.ConsensusState,
	header *ibctmtypes.Header, currentTimestamp time.Time,
) error {
	if err := checkTrustedHeader(header, consState); err != nil {
		return err
	}

	// UpdateClient only accepts updates with a header at the same revision
	// as the trusted consensus state
	if header.GetHeight().GetRevisionNumber() != header.TrustedHeight.RevisionNumber {
		return sdkerrors.Wrapf(
			ibctmtypes.ErrInvalidHeaderHeight,
			"header height revision %d does not match trusted header revision %d",
			header.GetHeight().GetRevisionNumber(), header.TrustedHeight.RevisionNumber,
		)
	}

	tmTrustedValidators, err := tmtypes.ValidatorSetFromProto(header.TrustedValidators)
	if err != nil {
		return sdkerrors.Wrap(err, "trusted validator set in not tendermint validator set type")
	}

	tmSignedHeader, err := tmtypes.SignedHeaderFromProto(header.SignedHeader)
	if err != nil {
		return sdkerrors.Wrap(err, "signed header in not tendermint signed header type")
	}

	tmValidatorSet, err := tmtypes.ValidatorSetFromProto(header.ValidatorSet)
	if err != nil {
		return sdkerrors.Wrap(err, "validator set in not tendermint validator set type")
	}

	// assert header height is newer than consensus state
	if header.GetHeight().LTE(header.TrustedHeight) {
		return sdkerrors.Wrapf(
			types.ErrInvalidHeader,
			"header height ≤ consensus state height (%s ≤ %s)", header.GetHeight(), header.TrustedHeight,
		)
	}

	chainID := clientState.GetChainID()
	// If chainID is in revision format, then set revision number of chainID with the revision number
	// of the header we are verifying
	// This is useful if the update is at a previous revision rather than an update to the latest revision
	// of the client.
	// The chainID must be set correctly for the previous revision before attempting verification.
	// Updates for previous revisions are not supported if the chainID is not in revision format.
	if types.IsRevisionFormat(chainID) {
		chainID, _ = types.SetRevisionNumber(chainID, header.GetHeight().GetRevisionNumber())
	}

	// Construct a trusted header using the fields in consensus state
	// Only Height, Time, and NextValidatorsHash are necessary for verification
	trustedHeader := tmtypes.Header{
		ChainID:            chainID,
		Height:             int64(header.TrustedHeight.RevisionHeight),
		Time:               consState.Timestamp,
		NextValidatorsHash: consState.NextValidatorsHash,
	}
	signedHeader := tmtypes.SignedHeader{
		Header: &trustedHeader,
	}

	// Verify next header with the passed-in trustedVals
	// - asserts trusting period not passed
	// - assert header timestamp is not past the trusting period
	// - assert header timestamp is past latest stored consensus state timestamp
	// - assert that a TrustLevel proportion of TrustedValidators signed new Commit
	err = light.Verify(
		&signedHeader,
		tmTrustedValidators, tmSignedHeader, tmValidatorSet,
		clientState.TrustingPeriod, currentTimestamp, clientState.MaxClockDrift, clientState.TrustLevel.ToTendermint(),
	)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to verify header")
	}
	return nil
}

// checkTrustedHeader checks that consensus state matches trusted fields of Header
// (copied from https://github.com/cosmos/ibc-go/blob/v5.0.0/modules/light-clients/07-tendermint/types/update.go#L148-L166)
func checkTrustedHeader(header *ibctmtypes.Header, consState *ibctmtypes.ConsensusState) error {
	tmTrustedValidators, err := tmtypes.ValidatorSetFromProto(header.TrustedValidators)
	if err != nil {
		return sdkerrors.Wrap(err, "trusted validator set in not tendermint validator set type")
	}

	// assert that trustedVals is NextValidators of last trusted header
	// to do this, we check that trustedVals.Hash() == consState.NextValidatorsHash
	tvalHash := tmTrustedValidators.Hash()
	if !bytes.Equal(consState.NextValidatorsHash, tvalHash) {
		return sdkerrors.Wrapf(
			ibctmtypes.ErrInvalidValidatorSet,
			"trusted validators %s, does not hash to latest trusted validators. Expected: %X, got: %X",
			header.TrustedValidators, consState.NextValidatorsHash, tvalHash,
		)
	}
	return nil
}
