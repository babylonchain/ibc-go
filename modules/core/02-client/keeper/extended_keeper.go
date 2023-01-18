package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	metrics "github.com/armon/go-metrics"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v5/modules/core/exported"
	"github.com/tendermint/tendermint/crypto/tmhash"
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

func (ek ExtendedKeeper) UpdateClient(ctx sdk.Context, clientID string, clientMsg exported.ClientMessage) error {
	clientState, found := ek.GetClientState(ctx, clientID)
	if !found {
		return sdkerrors.Wrapf(types.ErrClientNotFound, "cannot update client with ID %s", clientID)
	}

	clientStore := ek.ClientStore(ctx, clientID)

	if status := clientState.Status(ctx, clientStore, ek.cdc); status != exported.Active {
		return sdkerrors.Wrapf(types.ErrClientNotActive, "cannot update client (%s) with status %s", clientID, status)
	}

	if err := clientState.VerifyClientMessage(ctx, ek.cdc, clientStore, clientMsg); err != nil {
		return err
	}

	foundMisbehaviour := clientState.CheckForMisbehaviour(ctx, ek.cdc, clientStore, clientMsg)

	header, isHeader := clientMsg.(exported.LCHeader)

	// First difference in comparission to standard keeper update client
	if foundMisbehaviour && isHeader {
		// this is header and we foundMisbehaviour. Depending on lc client it can
		// mean many things in case of tendermint it means either invalid timestamp
		// or conflicting header. Do not update state and do not freeze light client
		ctx.Logger().Debug("received a header that has QC but is on a fork")
		txHash := tmhash.Sum(ctx.TxBytes())
		ek.AfterHeaderWithValidCommit(ctx, txHash, header, true)
		return nil
	}

	if foundMisbehaviour {
		clientState.UpdateStateOnMisbehaviour(ctx, ek.cdc, clientStore, clientMsg)

		ek.Logger(ctx).Info("client frozen due to misbehaviour", "client-id", clientID)

		defer telemetry.IncrCounterWithLabels(
			[]string{"ibc", "client", "misbehaviour"},
			1,
			[]metrics.Label{
				telemetry.NewLabel(types.LabelClientType, clientState.ClientType()),
				telemetry.NewLabel(types.LabelClientID, clientID),
				telemetry.NewLabel(types.LabelMsgType, "update"),
			},
		)

		EmitSubmitMisbehaviourEvent(ctx, clientID, clientState)

		return nil
	}

	// Second difference in comparission to standard keeper update client
	// this valid header without any misbehaviours time stamp it
	if isHeader {
		txHash := tmhash.Sum(ctx.TxBytes())
		ek.AfterHeaderWithValidCommit(ctx, txHash, header, false)
	}

	consensusHeights := clientState.UpdateState(ctx, ek.cdc, clientStore, clientMsg)

	ek.Logger(ctx).Info("client state updated", "client-id", clientID, "heights", consensusHeights)

	defer telemetry.IncrCounterWithLabels(
		[]string{"ibc", "client", "update"},
		1,
		[]metrics.Label{
			telemetry.NewLabel(types.LabelClientType, clientState.ClientType()),
			telemetry.NewLabel(types.LabelClientID, clientID),
			telemetry.NewLabel(types.LabelUpdateType, "msg"),
		},
	)

	// emitting events in the keeper emits for both begin block and handler client updates
	EmitUpdateClientEvent(ctx, clientID, clientState.ClientType(), consensusHeights, ek.cdc, clientMsg)

	return nil
}
