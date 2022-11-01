package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctmtypes "github.com/cosmos/ibc-go/v5/modules/light-clients/07-tendermint/types"
)

// ClientHooks defines the hook interface for client
type ClientHooks interface {
	AfterHeaderWithQC(ctx sdk.Context, header *ibctmtypes.Header)
}

// MultiClientHooks is a concrete implementation of ClientHooks
// It allows other modules to hook onto client ExtendedKeeper
var _ ClientHooks = &MultiClientHooks{}

type MultiClientHooks []ClientHooks

func NewMultiClientHooks(hooks ...ClientHooks) MultiClientHooks {
	return hooks
}

// invoke hooks in each keeper that hooks onto ExtendedKeeper
func (h MultiClientHooks) AfterHeaderWithQC(ctx sdk.Context, header *ibctmtypes.Header) {
	for i := range h {
		h[i].AfterHeaderWithQC(ctx, header)
	}
}

// ensure ExtendedKeeper implements ClientHooks interfaces
var _ ClientHooks = ExtendedKeeper{}

func (ek ExtendedKeeper) AfterHeaderWithQC(ctx sdk.Context, header *ibctmtypes.Header) {
	if ek.hooks != nil {
		ek.hooks.AfterHeaderWithQC(ctx, header)
	}
}
