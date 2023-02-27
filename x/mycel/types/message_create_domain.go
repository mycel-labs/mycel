package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateDomain = "create_domain"

var _ sdk.Msg = &MsgCreateDomain{}

func NewMsgCreateDomain(creator string, name string, parent string, registrationPeriodInYear uint64) *MsgCreateDomain {
	return &MsgCreateDomain{
		Creator:                  creator,
		Name:                     name,
		Parent:                   parent,
		RegistrationPeriodInYear: registrationPeriodInYear,
	}
}

func (msg *MsgCreateDomain) Route() string {
	return RouterKey
}

func (msg *MsgCreateDomain) Type() string {
	return TypeMsgCreateDomain
}

func (msg *MsgCreateDomain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDomain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDomain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
