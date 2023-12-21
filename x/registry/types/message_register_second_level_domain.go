package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterSecondLevelDomain = "register_domain"

var _ sdk.Msg = &MsgRegisterSecondLevelDomain{}

func NewMsgRegisterSecondLevelDomain(creator string, name string, parent string, registrationPeriodInYear uint64) *MsgRegisterSecondLevelDomain {
	return &MsgRegisterSecondLevelDomain{
		Creator:                  creator,
		Name:                     name,
		Parent:                   parent,
		RegistrationPeriodInYear: registrationPeriodInYear,
	}
}

func (msg *MsgRegisterSecondLevelDomain) Route() string {
	return RouterKey
}

func (msg *MsgRegisterSecondLevelDomain) Type() string {
	return TypeMsgRegisterSecondLevelDomain
}

func (msg *MsgRegisterSecondLevelDomain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterSecondLevelDomain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterSecondLevelDomain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
