package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// Define constants for your module's message types and routes
	TypeMsgDelegate = "delegate"
)

type MsgGridnodeDelegate struct {
	DelegatorAddress sdk.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	Amount           int64          `json:"amount" yaml:"amount"`
}

type MsgGridnodeUndelegate struct {
	DelegatorAddress sdk.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	Amount           int64          `json:"amount" yaml:"amount"`
}

// ProtoMessage implements types.Msg.
func (MsgGridnodeDelegate) ProtoMessage() {
	panic("unimplemented")
}

// Reset implements types.Msg.
func (MsgGridnodeDelegate) Reset() {
	panic("unimplemented")
}

// String implements types.Msg.
func (MsgGridnodeDelegate) String() string {
	panic("unimplemented")
}

func (MsgGridnodeUndelegate) ProtoMessage() {
	panic("unimplemented")
}

// Reset implements types.Msg.
func (MsgGridnodeUndelegate) Reset() {
	panic("unimplemented")
}

// String implements types.Msg.
func (MsgGridnodeUndelegate) String() string {
	panic("unimplemented")
}

func NewMsgDelegateGridnode(delegatorAddress sdk.AccAddress, amount int64) *MsgGridnodeDelegate {
	return &MsgGridnodeDelegate{
		DelegatorAddress: delegatorAddress,
		Amount:           amount,
	}
}

func NewMsgUndelegateGridnode(delegatorAddress sdk.AccAddress, amount int64) *MsgGridnodeUndelegate {
	return &MsgGridnodeUndelegate{
		DelegatorAddress: delegatorAddress,
		Amount:           amount,
	}
}

// Implementing the sdk.Msg interface

func (msg *MsgGridnodeDelegate) Route() string {
	return RouterKey
}

func (msg *MsgGridnodeDelegate) Type() string {
	return TypeMsgDelegate
}

func (msg MsgGridnodeDelegate) ValidateBasic() error {
	fmt.Println("Delegator Address:", msg.DelegatorAddress)
	if msg.DelegatorAddress.Empty() {
		return sdkerrors.ErrInvalidAddress
	}
	if msg.Amount <= 0 {
		return sdkerrors.ErrInvalidCoins
	}
	return nil
}

func (msg MsgGridnodeDelegate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgGridnodeDelegate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.DelegatorAddress}
}

func (msg *MsgGridnodeUndelegate) Route() string {
	return RouterKey
}

func (msg *MsgGridnodeUndelegate) Type() string {
	return "undelegate"
}

func (msg MsgGridnodeUndelegate) ValidateBasic() error {
	fmt.Println("Delegator Address:", msg.DelegatorAddress)
	if msg.DelegatorAddress.Empty() {
		return sdkerrors.ErrInvalidAddress
	}
	if msg.Amount <= 0 {
		return sdkerrors.ErrInvalidCoins
	}
	return nil
}

func (msg MsgGridnodeUndelegate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgGridnodeUndelegate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.DelegatorAddress}
}
