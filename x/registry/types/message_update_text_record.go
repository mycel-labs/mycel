package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateTextRecord = "update_text_record"

var _ sdk.Msg = &MsgUpdateTextRecord{}

func NewMsgUpdateTextRecord(creator string, name string, parent string, key string, value string) *MsgUpdateTextRecord {
	return &MsgUpdateTextRecord{
		Creator: creator,
		Name:    name,
		Parent:  parent,
		Key:     key,
		Value:   value,
	}
}

func (msg *MsgUpdateTextRecord) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTextRecord) Type() string {
	return TypeMsgUpdateTextRecord
}

func (msg *MsgUpdateTextRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTextRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTextRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
