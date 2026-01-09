package db

import (
	"context"
	"testing"

	"github.com/Umesh0910/simplebank/db/util"
	"github.com/stretchr/testify/require"
)

func createTestAccount(t *testing.T) Account {
	ctx := context.Background()
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Currency: util.RandomCurrency(),
		Balance:  util.RandomBalance(),
	}

	resp, err := testQueries.CreateAccount(ctx, args)
	require.NoError(t, err)
	require.Equal(t, resp.Owner, args.Owner)
	require.Equal(t, resp.Currency, args.Currency)
	require.Equal(t, resp.Balance, args.Balance)
	require.NotZero(t, resp.ID, resp.CreatedAt)

	return resp
}

func TestCreateAccount(t *testing.T) {
	createTestAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createTestAccount(t)

	ctx := context.Background()
	resp, err := testQueries.GetAccount(ctx, account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, resp)
	require.Equal(t, account.ID, resp.ID)
	require.Equal(t, account.Balance, resp.Balance)
	require.Equal(t, account.Currency, resp.Currency)
	require.Equal(t, account.Owner, resp.Owner)
}

func TestUpdateAccount(t *testing.T) {
	account := createTestAccount(t)
	args := UpdateAccountParams{
		ID:       account.ID,
		Owner:    util.RandomOwner(),
		Currency: util.RandomCurrency(),
		Balance:  util.RandomBalance(),
	}
	ctx := context.Background()
	_,err := testQueries.UpdateAccount(ctx, args)
	require.NoError(t, err)
}

func TestDeleteAccount(t *testing.T) {
	account := createTestAccount(t)

	ctx := context.Background()
	err := testQueries.DeleteAccount(ctx, account.ID)
	require.NoError(t, err)
}
