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

type MsgDelegate struct {
	DelegatorAddress sdk.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	Amount           int64          `json:"amount" yaml:"amount"`
}

type MsgUndelegate struct {
	DelegatorAddress sdk.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	Amount           int64          `json:"amount" yaml:"amount"`
}

// ProtoMessage implements types.Msg.
func (MsgDelegate) ProtoMessage() {
	panic("unimplemented")
}

// Reset implements types.Msg.
func (MsgDelegate) Reset() {
	panic("unimplemented")
}

// String implements types.Msg.
func (MsgDelegate) String() string {
	panic("unimplemented")
}

func (MsgUndelegate) ProtoMessage() {
	panic("unimplemented")
}

// Reset implements types.Msg.
func (MsgUndelegate) Reset() {
	panic("unimplemented")
}

// String implements types.Msg.
func (MsgUndelegate) String() string {
	panic("unimplemented")
}

func NewMsgDelegateGridnode(delegatorAddress sdk.AccAddress, amount int64) *MsgDelegate {
	return &MsgDelegate{
		DelegatorAddress: delegatorAddress,
		Amount:           amount,
	}
}

func NewMsgUndelegateGridnode(delegatorAddress sdk.AccAddress, amount int64) *MsgUndelegate {
	return &MsgUndelegate{
		DelegatorAddress: delegatorAddress,
		Amount:           amount,
	}
}

// Implementing the sdk.Msg interface

func (msg *MsgDelegate) Route() string {
	return RouterKey
}

func (msg *MsgDelegate) Type() string {
	return TypeMsgDelegate
}

func (msg MsgDelegate) ValidateBasic() error {
	fmt.Println("Delegator Address:", msg.DelegatorAddress)
	if msg.DelegatorAddress.Empty() {
		return sdkerrors.ErrInvalidAddress
	}
	if msg.Amount <= 0 {
		return sdkerrors.ErrInvalidCoins
	}
	return nil
}

func (msg MsgDelegate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgDelegate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.DelegatorAddress}
}

func (msg *MsgUndelegate) Route() string {
	return RouterKey
}

func (msg *MsgUndelegate) Type() string {
	return "undelegate"
}

func (msg MsgUndelegate) ValidateBasic() error {
	fmt.Println("Delegator Address:", msg.DelegatorAddress)
	if msg.DelegatorAddress.Empty() {
		return sdkerrors.ErrInvalidAddress
	}
	if msg.Amount <= 0 {
		return sdkerrors.ErrInvalidCoins
	}
	return nil
}

func (msg MsgUndelegate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgUndelegate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.DelegatorAddress}
}
