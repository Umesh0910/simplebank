package db

import (
	"context"
	"testing"

	"github.com/Umesh0910/simplebank/db/util"
	"github.com/stretchr/testify/require"
)

func createTestTransfer(t *testing.T, account Account, account2 Account) Transfer {
	args := CreateTransferParams{
		FromAccountID: account.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomBalance(),
	}
	ctx := context.Background()
	resp, err := testQueries.CreateTransfer(ctx, args)

	require.NoError(t, err)
	require.Equal(t, args.FromAccountID, resp.FromAccountID)
	require.Equal(t, args.ToAccountID, resp.ToAccountID)
	require.Equal(t, args.Amount, resp.Amount)

	require.NotZero(t, resp.ID, resp.CreatedAt)

	return resp
}

func TestCreateTransfer(t *testing.T) {
	account := createTestAccount(t)
	account2 := createTestAccount(t)
	createTestTransfer(t, account, account2)
}

func TestGetTransfer(t *testing.T) {
	account := createTestAccount(t)
	account2 := createTestAccount(t)
	Transfer := createTestTransfer(t, account, account2)

	ctx := context.Background()
	resp, err := testQueries.GetTransfer(ctx, Transfer.ID)

	require.NoError(t, err)
	require.Equal(t, Transfer.ID, resp.ID)
	require.Equal(t, Transfer.FromAccountID, resp.FromAccountID)
	require.Equal(t, Transfer.ToAccountID, resp.ToAccountID)
	require.Equal(t, Transfer.Amount, resp.Amount)
	require.NotEmpty(t, resp.CreatedAt, resp.ID)
}

func TestListTransfers(t *testing.T) {
	account := createTestAccount(t)
	account2 := createTestAccount(t)
	for i := 0; i < 10; i++ {
		createTestTransfer(t, account, account2)
		createTestTransfer(t, account2, account)
	}

	args := ListTransfersParams{
		FromAccountID: account.ID,
		ToAccountID: account2.ID,
		Limit:  5,
		Offset: 5,
	}

	ctx := context.Background()
	resp, err := testQueries.ListTransfers(ctx, args)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
	require.Len(t, resp, 5)

	for _, transfer:= range resp {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account.ID || transfer.ToAccountID == account.ID)
	}
}

func TestDeleteTransfer(t *testing.T) {
	account := createTestAccount(t)
	account2 := createTestAccount(t)
	Transfer := createTestTransfer(t, account, account2)

	ctx := context.Background()
	err := testQueries.DeleteTransfer(ctx, Transfer.ID)

	require.NoError(t, err)
}

func TestUpdateTransfer(t *testing.T){
	account := createTestAccount(t)
	account2 := createTestAccount(t)
	createTestTransfer(t, account, account2)

	args := UpdateTransferParams{
		FromAccountID: account.ID,
		ToAccountID: account2.ID,
		Amount: util.RandomBalance(),
	}

	ctx := context.Background()
	err := testQueries.UpdateTransfer(ctx, args)

	require.NoError(t, err)
}