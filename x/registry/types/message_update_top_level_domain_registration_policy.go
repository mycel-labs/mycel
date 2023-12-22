package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateTopLevelDomainRegistrationPolicy = "update_top_level_domain_registration_policy"

var _ sdk.Msg = &MsgUpdateTopLevelDomainRegistrationPolicy{}

func NewMsgUpdateTopLevelDomainRegistrationPolicy(creator string, name string, registrationPolicy string) *MsgUpdateTopLevelDomainRegistrationPolicy {
	return &MsgUpdateTopLevelDomainRegistrationPolicy{
		Creator:            creator,
		Name:               name,
		RegistrationPolicy: registrationPolicy,
	}
}

func (msg *MsgUpdateTopLevelDomainRegistrationPolicy) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTopLevelDomainRegistrationPolicy) Type() string {
	return TypeMsgUpdateTopLevelDomainRegistrationPolicy
}

func (msg *MsgUpdateTopLevelDomainRegistrationPolicy) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTopLevelDomainRegistrationPolicy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTopLevelDomainRegistrationPolicy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
