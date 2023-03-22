package keeper_test

import (
	"context"
	"fmt"
	"testing"

	"mycel/testutil"
	"mycel/x/registry/keeper"
	"mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

type TestMsgUpdateWalletRecord struct {
	MsgUpdateWalletRecord types.MsgUpdateWalletRecord
	IsInvalidDomain       bool
	IsInvalidOwner        bool
}

func GetValidMsgUpdateWalletRecords() []TestMsgUpdateWalletRecord {
	return []TestMsgUpdateWalletRecord{
		{
			MsgUpdateWalletRecord: types.MsgUpdateWalletRecord{
				Creator:          testutil.Alice,
				Name:             "foo",
				Parent:           "cel",
				WalletRecordType: "ETHEREUM_MAINNET",
				Value:            "0x1234567890123456789012345678901234567890",
			},
		},
		{
			MsgUpdateWalletRecord: types.MsgUpdateWalletRecord{
				Creator:          testutil.Alice,
				Name:             "foo",
				Parent:           "cel",
				WalletRecordType: "ETHEREUM_MAINNET",
				Value:            "0x1234567890123456789012345678901234567891",
			},
		},
		{
			MsgUpdateWalletRecord: types.MsgUpdateWalletRecord{
				Creator:          testutil.Alice,
				Name:             "foo",
				Parent:           "cel",
				WalletRecordType: "ETHEREUM_GOERLI",
				Value:            "0x1234567890123456789012345678901234567890",
			},
		},
		{
			MsgUpdateWalletRecord: types.MsgUpdateWalletRecord{
				Creator:          testutil.Alice,
				Name:             "foo",
				Parent:           "cel",
				WalletRecordType: "POLYGON_MAINNET",
				Value:            "0x1234567890123456789012345678901234567890",
			},
		},
		{
			MsgUpdateWalletRecord: types.MsgUpdateWalletRecord{
				Creator:          testutil.Alice,
				Name:             "foo",
				Parent:           "cel",
				WalletRecordType: "POLYGON_MUMBAI",
				Value:            "0x1234567890123456789012345678901234567890",
			},
		},
	}
}
func GetInvalidMsgUpdateWalletRecords() []TestMsgUpdateWalletRecord {
	return []TestMsgUpdateWalletRecord{
		{
			MsgUpdateWalletRecord: types.MsgUpdateWalletRecord{
				Creator:          testutil.Alice,
				Name:             "hoge",
				Parent:           "fuga",
				WalletRecordType: "ETHEREUM_MAINNET",
				Value:            "0x1234567890123456789012345678901234567890",
			},
			IsInvalidDomain: true,
		},
		{
			MsgUpdateWalletRecord: types.MsgUpdateWalletRecord{
				Creator:          testutil.Bob,
				Name:             "foo",
				Parent:           "cel",
				WalletRecordType: "ETHEREUM_MAINNET",
				Value:            "0x1234567890123456789012345678901234567890",
			},
			IsInvalidOwner: true,
		},
	}
}

func setupMsgServerUpdateRecord(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	msgServer, keeper, context, ctrl, escrow := setupMsgServerWithMock(t)
	defer ctrl.Finish()
	escrow.ExpectAny(context)

	// Register domain
	domain := GetMsgRegisterDomain()
	_, err := msgServer.RegisterDomain(context, domain)
	require.Nil(t, err)

	return msgServer, keeper, context
}

func TestUpdateWalletRecordSuccess(t *testing.T) {
	msgServer, keeper, context := setupMsgServerUpdateRecord(t)

	for i, record := range GetValidMsgUpdateWalletRecords() {
		// Update wallet record
		_, err := msgServer.UpdateWalletRecord(context, &record.MsgUpdateWalletRecord)
		require.Nil(t, err)

		domain := GetMsgRegisterDomain()

		// Check if wallet record is UpdateWalletRecord
		ctx := sdk.UnwrapSDKContext(context)
		require.NotNil(t, ctx)
		res, _ := keeper.GetDomain(ctx, domain.Name, domain.Parent)
		require.Equal(t, record.MsgUpdateWalletRecord.Value, res.WalletRecords[record.MsgUpdateWalletRecord.WalletRecordType].Value)

		// Event check
		events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
		require.Len(t, events, 2)
		require.EqualValues(t,
			[]sdk.Attribute{
				{Key: types.AttributeUpdateWalletRecordEventDomainName, Value: domain.Name},
				{Key: types.AttributeUpdateWalletRecordEventDomainParent, Value: domain.Parent},
				{Key: types.AttributeUpdateWalletRecordEventWalletRecordType, Value: record.MsgUpdateWalletRecord.WalletRecordType},
				{Key: types.AttributeUpdateWalletRecordEventValue, Value: record.MsgUpdateWalletRecord.Value},
			},
			events[1].Attributes[i*4:])

	}

}

func TestUpdateWalletRecordDomainNotFoundFailure(t *testing.T) {
	msgServer, _, context := setupMsgServerUpdateRecord(t)

	for _, record := range GetInvalidMsgUpdateWalletRecords() {
		if record.IsInvalidDomain {
			_, err := msgServer.UpdateWalletRecord(context, &record.MsgUpdateWalletRecord)
			require.EqualError(t, err, fmt.Sprintf("domain not found: %s.%s", record.MsgUpdateWalletRecord.Name, record.MsgUpdateWalletRecord.Parent))
		}
	}
}

func TestUpdateWalletRecordNotOwnerFailure(t *testing.T) {
	msgServer, _, context := setupMsgServerUpdateRecord(t)

	for _, record := range GetInvalidMsgUpdateWalletRecords() {
		if record.IsInvalidOwner {
			_, err := msgServer.UpdateWalletRecord(context, &record.MsgUpdateWalletRecord)
			require.EqualError(t, err, fmt.Sprintf("domain not owned by msg creator: %s.%s", record.MsgUpdateWalletRecord.Name, record.MsgUpdateWalletRecord.Parent))
		}
	}
}
