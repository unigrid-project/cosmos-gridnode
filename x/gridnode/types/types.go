package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Gridnode struct {
	ID    string         // This is the gridnode-id from Hedgehog
	Owner sdk.AccAddress // This is the owner of the gridnode
}

type KeyInfo struct {
	Keys      []string `json:"keys"`
	Timestamp int64    `json:"timestamp"`
	UniqueId  string   `json:"unique_id"`
	Status    string   `json:"status"`
}

type AccountKeys struct {
	Account   string   `json:"account_address"`
	Keys      []string `json:"keys"`
	Timestamp int64    `json:"timestamp"`
	UniqueId  string   `json:"unique_id"`
	Status    string   `json:"status"`
}

// type UnbondingEntry struct {
// 	Delegator      sdk.AccAddress `json:"delegator" yaml:"delegator"`
// 	Amount         math.Int       `json:"amount" yaml:"amount"`
// 	CompletionTime time.Time      `json:"completion_time" yaml:"completion_time"`
// }

const (
	EventTypeDelegate          = "delegate"
	EventTypeUndelegate        = "undelegate"
	AttributeKeyDelegator      = "delegator"
	AttributeKeyAmount         = "amount"
	EventTypeCompleteUnbond    = "complete_unbond"
	EventTypeUnbond            = "unbond"
	AttributeKeyCompletionTime = "completion_time"
)

var ErrInsufficientFunds = errors.Register(ModuleName, 1100, "insufficient funds")
var ErrAmountExceedsDelagation = errors.Register(ModuleName, 1101, "amount exceeds delegated amount")
var ErrUnknownRequest = errors.Register(ModuleName, 101, "unknown request")
var ErrOverUnbond = errors.Register(ModuleName, 1102, "attempting to unbond more than the delegated amount")
