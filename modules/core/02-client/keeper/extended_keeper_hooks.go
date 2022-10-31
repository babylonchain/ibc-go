package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctmtypes "github.com/cosmos/ibc-go/v5/modules/light-clients/07-tendermint/types"
)

// ClientHooks defines the hook interface for client
type ClientHooks interface {
	AfterHeaderWithValidCommit(ctx sdk.Context, txHash []byte, header *ibctmtypes.Header, isOnFork bool)
}

// MultiClientHooks is a concrete implementation of ClientHooks
// It allows other modules to hook onto client ExtendedKeeper
var _ ClientHooks = &MultiClientHooks{}

type MultiClientHooks []ClientHooks

func NewMultiClientHooks(hooks ...ClientHooks) MultiClientHooks {
	return hooks
}

// invoke hooks in each keeper that hooks onto ExtendedKeeper
func (h MultiClientHooks) AfterHeaderWithValidCommit(ctx sdk.Context, txHash []byte, header *ibctmtypes.Header, isOnFork bool) {
	for i := range h {
		h[i].AfterHeaderWithValidCommit(ctx, txHash, header, isOnFork)
	}
}

// ensure ExtendedKeeper implements ClientHooks interfaces
var _ ClientHooks = ExtendedKeeper{}

func (ek ExtendedKeeper) AfterHeaderWithValidCommit(ctx sdk.Context, txHash []byte, header *ibctmtypes.Header, isOnFork bool) {
	if ek.hooks != nil {
		ek.hooks.AfterHeaderWithValidCommit(ctx, txHash, header, isOnFork)
	}
}
