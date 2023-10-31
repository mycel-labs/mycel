package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mycel-domain/mycel/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgRegisterSecondLevelDomain_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRegisterSecondLevelDomain
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRegisterSecondLevelDomain{
				Creator:                  "invalid_address",
				Name:                     "foo",
				Parent:                   "cel",
				RegistrationPeriodInYear: 1,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRegisterSecondLevelDomain{
				Creator:                  sample.AccAddress(),
				Name:                     "foo",
				Parent:                   "cel",
				RegistrationPeriodInYear: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
