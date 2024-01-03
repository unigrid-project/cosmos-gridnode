package keeper

import (
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/types"
)

var _ types.QueryServer = Keeper{}
