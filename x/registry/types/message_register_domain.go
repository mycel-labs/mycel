package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterDomain = "register_domain"

var _ sdk.Msg = &MsgRegisterDomain{}

func NewMsgRegisterDomain(creator string, name string, parent string, registrationPeriodInYear uint64) *MsgRegisterDomain {
	return &MsgRegisterDomain{
		Creator:                  creator,
		Name:                     name,
		Parent:                   parent,
		RegistrationPeriodInYear: registrationPeriodInYear,
	}
}

func (msg *MsgRegisterDomain) Route() string {
	return RouterKey
}

func (msg *MsgRegisterDomain) Type() string {
	return TypeMsgRegisterDomain
}

func (msg *MsgRegisterDomain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterDomain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterDomain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
