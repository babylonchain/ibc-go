package v7

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	clientexported "github.com/cosmos/ibc-go/v7/modules/core/02-client/exported"
)

// MigrateLocalhostClient initialises the 09-localhost client state and sets it in state.
func MigrateLocalhostClient(ctx sdk.Context, clientKeeper clientexported.ClientKeeper) error {
	return clientKeeper.CreateLocalhostClient(ctx)
}
