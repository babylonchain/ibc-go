package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
	ibctesting "github.com/cosmos/ibc-go/v5/testing"
)

func (suite *KeeperTestSuite) TestQueryInterchainAccount() {
	var msg *types.QueryInterchainAccountRequest

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success",
			func() {},
			true,
		},
		{
			"empty request",
			func() {
				msg = nil
			},
			false,
		},
		{
			"empty owner address",
			func() {
				msg.Owner = ""
			},
			false,
		},
		{
			"invalid connection, account address not found",
			func() {
				msg.ConnectionId = "invalid-connection-id"
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			path := NewICAPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupConnections(path)

			err := SetupICAPath(path, ibctesting.TestAccAddress)
			suite.Require().NoError(err)

			msg = &types.QueryInterchainAccountRequest{
				ConnectionId: ibctesting.FirstConnectionID,
				Owner:        ibctesting.TestAccAddress,
			}

			tc.malleate()

			res, err := suite.chainA.GetSimApp().ICAControllerKeeper.InterchainAccount(sdk.WrapSDKContext(suite.chainA.GetContext()), msg)

			if tc.expPass {
				expAddress := icatypes.GenerateAddress(suite.chainA.GetSimApp().AccountKeeper.GetModuleAddress(icatypes.ModuleName), path.EndpointB.ConnectionID, path.EndpointA.ChannelConfig.PortID)

				suite.Require().NoError(err)
				suite.Require().Equal(expAddress.String(), res.Address)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryParams() {
	ctx := sdk.WrapSDKContext(suite.chainA.GetContext())
	expParams := types.DefaultParams()
	res, _ := suite.chainA.GetSimApp().ICAControllerKeeper.Params(ctx, &types.QueryParamsRequest{})
	suite.Require().Equal(&expParams, res.Params)
}
