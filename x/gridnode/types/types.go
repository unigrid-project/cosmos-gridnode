package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Gridnode struct {
	ID    string         // This is the gridnode-id from Hedgehog
	Owner sdk.AccAddress // This is the owner of the gridnode
}

const (
	EventTypeDelegate     = "delegate"
	AttributeKeyDelegator = "delegator"
	AttributeKeyAmount    = "amount"
)

var ErrInsufficientFunds = errors.Register(ModuleName, 1100, "insufficient funds")
var ErrUnknownRequest = errors.Register(ModuleName, 101, "unknown request")

func IsGridnode(voterAddr sdk.AccAddress) bool {
	res := true

	if res {
		return true
	}

	return false
}
