package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ MsgDelegate = &MsgDelegate{}
)

// NewMsgDelegate creates a new MsgDelegate instance.
func NewMsgDelegate(delAddr string, amount sdk.Coin) *MsgDelegate {
	return &MsgDelegate{
		DelegatorAddress: delAddr,
		Amount:           amount,
	}
}
