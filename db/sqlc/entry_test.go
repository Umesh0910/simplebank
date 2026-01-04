package db

import (
	"context"
	"testing"

	"github.com/Umesh0910/simplebank/db/util"
	"github.com/stretchr/testify/require"
)

func createTestEntry(t *testing.T, account Account) Entry {
	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomBalance(),
	}
	ctx := context.Background()
	resp, err := testQueries.CreateEntry(ctx, args)

	require.NoError(t, err)
	require.Equal(t, args.AccountID, resp.AccountID)
	require.Equal(t, args.Amount, resp.Amount)

	require.NotZero(t, resp.ID, resp.CreatedAt)

	return resp
}

func TestCreateEntry(t *testing.T) {
	account := createTestAccount(t)
	createTestEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createTestAccount(t)
	entry := createTestEntry(t, account)

	ctx := context.Background()
	resp, err := testQueries.GetEntry(ctx, entry.ID)

	require.NoError(t, err)
	require.Equal(t, entry.ID, resp.ID)
	require.Equal(t, entry.Amount, resp.Amount)
	require.NotEmpty(t, resp.CreatedAt, resp.ID)
}

func TestListEntries(t *testing.T) {
	account := createTestAccount(t)
	for i := 0; i < 10; i++ {
		createTestEntry(t, account)
	}

	args := ListEntriesParams{
		AccountID: account.ID,
		Limit:  5,
		Offset: 5,
	}

	ctx := context.Background()
	resp, err := testQueries.ListEntries(ctx, args)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
	require.Len(t, resp, 5)

	for _, entry := range resp {
		require.NotEmpty(t, entry)
		require.Equal(t, entry.AccountID, args.AccountID)
	}
}

func TestDeleteEntry(t *testing.T) {
	account := createTestAccount(t)
	entry := createTestEntry(t, account)

	ctx := context.Background()
	err := testQueries.DeleteEntry(ctx, entry.ID)

	require.NoError(t, err)
}

func TestUpdateEntry(t *testing.T){
	account := createTestAccount(t)
	createTestEntry(t, account)

	args := UpdateEntryParams{
		AccountID: account.ID,
		Amount: util.RandomBalance(),
	}

	ctx := context.Background()
	err := testQueries.UpdateEntry(ctx, args)

	require.NoError(t, err)
}