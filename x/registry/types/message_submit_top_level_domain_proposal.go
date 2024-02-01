package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitTopLevelDomainProposal = "submit_top_level_domain_proposal"

var _ sdk.Msg = &MsgSubmitTopLevelDomainProposal{}

func NewMsgSubmitTopLevelDomainProposal(creator string, name string, registrationPeriodInYear uint64) *MsgSubmitTopLevelDomainProposal {
	return &MsgSubmitTopLevelDomainProposal{
		Creator:                  creator,
		Name:                     name,
		RegistrationPeriodInYear: registrationPeriodInYear,
	}
}

func (msg *MsgSubmitTopLevelDomainProposal) Route() string {
	return RouterKey
}

func (msg *MsgSubmitTopLevelDomainProposal) Type() string {
	return TypeMsgSubmitTopLevelDomainProposal
}

func (msg *MsgSubmitTopLevelDomainProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitTopLevelDomainProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitTopLevelDomainProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
