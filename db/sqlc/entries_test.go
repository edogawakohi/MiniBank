package db

import (
	"context"
	"minibank/utils"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func randomCreateEntry(t *testing.T) Entry {

	//create account
	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, account.ID)
	require.Equal(t, entry.Amount, arg.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	randomCreateEntry(t)
}

func TestDeleteEntry(t *testing.T) {
	entry := randomCreateEntry(t)

	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntryAnAccount(context.Background(), entry.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Equal(t, int64(0), entry2.EntryID)
}

func TestDeleteEntryWithGetEntry(t *testing.T) {
	entry := randomCreateEntry(t)

	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, entry2)
}

func TestGetEntry(t *testing.T) {

	entry1 := randomCreateEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt.Time, entry2.CreatedAt.Time, time.Second)

}

func TestGetEntryAnAccount(t *testing.T) {
	entry1 := randomCreateEntry(t)

	entry2, err := testQueries.GetEntryAnAccount(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.ID, entry2.EntryID)

	require.NotEmpty(t, entry2.Owner)
	require.NotZero(t, entry2.Balance)

}

func TestListEntriesByAccount(t *testing.T) {

	account := createRandomAccount(t)

	for i := 0; i <= 10; i++ {
		arg := CreateEntryParams{
			AccountID: account.ID,
			Amount:    utils.RandomMoney(),
		}
		_, err := testQueries.CreateEntry(context.Background(), arg)
		require.NoError(t, err)
	}

	args := ListEntriesByAccountParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    0,
	}

	entries, err := testQueries.ListEntriesByAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, account.ID, entry.AccountID)
	}
}

func TestUpdateEntry(t *testing.T) {
	entry1 := randomCreateEntry(t)

	arg := UpdateEntryAmountParams{
		ID:     entry1.ID,
		Amount: utils.RandomMoney(),
	}

	entry2, err := testQueries.UpdateEntryAmount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, arg.Amount, entry2.Amount)
	require.Equal(t, arg.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)

}
