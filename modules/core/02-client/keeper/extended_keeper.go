package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v5/modules/core/exported"
	ibctmtypes "github.com/cosmos/ibc-go/v5/modules/light-clients/07-tendermint/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

// ExtendedKeeper is same as the original Keeper, except that
// - it provides hooks for notifying other modules on received headers
// - it applies different verification rules on received headers
//   (notably, choosing forks rather than freezing clients upon conflicted headers with QCS)
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

func (ek ExtendedKeeper) UpdateClient(ctx sdk.Context, clientID string, header exported.Header) error {
	// header can be nil in the original IBC-Go design, e.g., in `BeginBlocker`
	// Our logic only applies when header is not nil
	if header != nil {
		// asserting
		// copied from `modules/light-clients/07-tendermint/types/update.go#L57`
		tmHeader, ok := header.(*ibctmtypes.Header)
		if !ok {
			return sdkerrors.Wrapf(
				types.ErrInvalidHeader, "expected type %T, got %T", &ibctmtypes.Header{}, header,
			)
		}

		// TODO: ensure the header has a valid QC

		// get hash of the tx that includes this header
		txHash := tmhash.Sum(ctx.TxBytes())

		// invoke hooks
		ek.AfterHeaderWithValidCommit(ctx, txHash, tmHeader)
	}

	// TODO: inject our own verification rules
	return ek.Keeper.UpdateClient(ctx, clientID, header)
}
