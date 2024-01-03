package gridnode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/unigrid-project/cosmos-gridnode/testutil/keeper"
	"github.com/unigrid-project/cosmos-gridnode/testutil/nullify"
	gridnode "github.com/unigrid-project/cosmos-gridnode/x/gridnode/module"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GridnodeKeeper(t)
	gridnode.InitGenesis(ctx, *k, genesisState)
	got := gridnode.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
