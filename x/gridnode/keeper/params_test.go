package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/unigrid-project/cosmos-gridnode/testutil/keeper"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.GridnodeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
