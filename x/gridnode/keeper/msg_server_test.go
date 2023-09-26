package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/unigrid-project/cosmos-sdk-gridnode/testutil/keeper"
	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/keeper"
	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.GridnodeKeeper(t)
	return keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
