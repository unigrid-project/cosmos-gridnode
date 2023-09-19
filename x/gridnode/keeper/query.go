package keeper

import (
	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/types"
)

var _ types.QueryServer = Keeper{}
