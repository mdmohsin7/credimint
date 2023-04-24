package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLiquid = "liquid"

var _ sdk.Msg = &MsgLiquid{}

func NewMsgLiquid(creator string, amount string, validator string) *MsgLiquid {
	return &MsgLiquid{
		Creator:   creator,
		Amount:    amount,
		Validator: validator,
	}
}

func (msg *MsgLiquid) Route() string {
	return RouterKey
}

func (msg *MsgLiquid) Type() string {
	return TypeMsgLiquid
}

func (msg *MsgLiquid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLiquid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLiquid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
