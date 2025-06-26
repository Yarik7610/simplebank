package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Yarik7610/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	created := createRandomAccount(t)
	got, err := testQueries.GetAccount(context.Background(), created.ID)

	require.NoError(t, err)
	require.NotEmpty(t, got)

	require.Equal(t, created.ID, got.ID)
	require.Equal(t, created.Owner, got.Owner)
	require.Equal(t, created.Balance, got.Balance)
	require.Equal(t, created.Currency, got.Currency)
	require.WithinDuration(t, created.CreatedAt, got.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	created := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      created.ID,
		Balance: utils.RandomMoney(),
	}

	updated, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updated)

	require.Equal(t, arg.Balance, updated.Balance)

	require.Equal(t, created.ID, updated.ID)
	require.Equal(t, created.Owner, updated.Owner)
	require.WithinDuration(t, created.CreatedAt, updated.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	created := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), created.ID)

	require.NoError(t, err)

	existingAccount, err := testQueries.GetAccount(context.Background(), created.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, existingAccount)
}

func TestListAccount(t *testing.T) {
	for range 10 {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 3,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
