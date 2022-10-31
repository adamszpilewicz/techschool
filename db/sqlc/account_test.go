package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"techschool/util"
	"testing"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance, account2.Balance)
}

func TestUpdateAccount(t *testing.T) {
	accountBefore := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      accountBefore.ID,
		Balance: util.RandomMoney(),
	}
	accountAfter, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Balance, accountAfter.Balance)
}

func TestDeleteAccount(t *testing.T) {
	accountBefore := createRandomAccount(t)
	require.NotEmpty(t, accountBefore)
	err := testQueries.DeleteAccount(context.Background(), accountBefore.ID)
	require.NoError(t, err)

	accountFind, err := testQueries.GetAccount(context.Background(), accountBefore.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accountFind)

}
