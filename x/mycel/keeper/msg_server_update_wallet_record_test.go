package keeper_test

import (
	"fmt"
	"testing"

	"mycel/testutil"
	"mycel/x/mycel/types"

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
				Name:             "poyo",
				Parent:           "ninniku",
				WalletRecordType: "ETHEREUM_MAINNET",
				Value:            "0x1234567890123456789012345678901234567890",
			},
		},
		{
			MsgUpdateWalletRecord: types.MsgUpdateWalletRecord{
				Creator:          testutil.Alice,
				Name:             "poyo",
				Parent:           "ninniku",
				WalletRecordType: "ETHEREUM_MAINNET",
				Value:            "0x1234567890123456789012345678901234567891",
			},
		},
		{
			MsgUpdateWalletRecord: types.MsgUpdateWalletRecord{
				Creator:          testutil.Alice,
				Name:             "poyo",
				Parent:           "ninniku",
				WalletRecordType: "ETHEREUM_GOERLI",
				Value:            "0x1234567890123456789012345678901234567890",
			},
		},
		{
			MsgUpdateWalletRecord: types.MsgUpdateWalletRecord{
				Creator:          testutil.Alice,
				Name:             "poyo",
				Parent:           "ninniku",
				WalletRecordType: "POLYGON_MAINNET",
				Value:            "0x1234567890123456789012345678901234567890",
			},
		},
		{
			MsgUpdateWalletRecord: types.MsgUpdateWalletRecord{
				Creator:          testutil.Alice,
				Name:             "poyo",
				Parent:           "ninniku",
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
				Name:             "poyo",
				Parent:           "ninniku",
				WalletRecordType: "ETHEREUM_MAINNET",
				Value:            "0x1234567890123456789012345678901234567890",
			},
			IsInvalidOwner: true,
		},
	}
}

func TestUpdateWalletRecordSuccess(t *testing.T) {
	msgServer, keeper, context := setupMsgServer(t)

	// Register domain
	domain := GetMsgRegisterDomain()
	_, err := msgServer.RegisterDomain(context, domain)
	require.Nil(t, err)

	for _, record := range GetValidMsgUpdateWalletRecords() {
		// Update wallet record
		_, err = msgServer.UpdateWalletRecord(context, &record.MsgUpdateWalletRecord)
		require.Nil(t, err)

		// Check if wallet record is UpdateWalletRecord
		ctx := sdk.UnwrapSDKContext(context)
		res, _ := keeper.GetDomain(ctx, domain.Name, domain.Parent)
		require.Equal(t, record.MsgUpdateWalletRecord.Value, res.WalletRecords[record.MsgUpdateWalletRecord.WalletRecordType].Value)
	}

}

func TestUpdateWalletRecordDomainNotFoundFailure(t *testing.T) {
	msgServer, _, context := setupMsgServer(t)

	// Register domain
	domain := GetMsgRegisterDomain()
	_, err := msgServer.RegisterDomain(context, domain)
	require.Nil(t, err)

	for _, record := range GetInvalidMsgUpdateWalletRecords() {
		if record.IsInvalidDomain {
			_, err := msgServer.UpdateWalletRecord(context, &record.MsgUpdateWalletRecord)
			require.EqualError(t, err, fmt.Sprintf("domain not found: %s.%s", record.MsgUpdateWalletRecord.Name, record.MsgUpdateWalletRecord.Parent))
		}
	}
}

func TestUpdateWalletRecordNotOwnerFailure(t *testing.T) {
	msgServer, _, context := setupMsgServer(t)

	// Register domain
	domain := GetMsgRegisterDomain()
	_, err := msgServer.RegisterDomain(context, domain)
	require.Nil(t, err)

	for _, record := range GetInvalidMsgUpdateWalletRecords() {
		if record.IsInvalidOwner {
			_, err := msgServer.UpdateWalletRecord(context, &record.MsgUpdateWalletRecord)
			require.EqualError(t, err, fmt.Sprintf("domain not owned by msg creator: %s.%s", record.MsgUpdateWalletRecord.Name, record.MsgUpdateWalletRecord.Parent))
		}
	}
}
