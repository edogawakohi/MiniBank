package db

import (
	"context"
	"minibank/utils"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func randomCreateTransfer(t *testing.T) Transfer {

	from := createRandomAccount(t)
	to := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: from.ID,
		ToAccountID:   to.ID,
		Amount:        utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.CreatedAt)
	require.NotZero(t, transfer.ID)

	return transfer
}

func randomCreateTransferWithAccount(t *testing.T, fromAccount, toAccount int64) Transfer {

	arg := CreateTransferParams{
		FromAccountID: fromAccount,
		ToAccountID:   toAccount,
		Amount:        utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.CreatedAt)
	require.NotZero(t, transfer.ID)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	randomCreateTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := randomCreateTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt.Time, transfer2.CreatedAt.Time, time.Second)

}

func TestGetTransferToFromAccount(t *testing.T) {

	//create 2 account
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	//create multiple transfer
	for i := 0; i < 5; i++ {
		randomCreateTransferWithAccount(t, fromAccount.ID, toAccount.ID)

	}

	arg := GetTransferToFromAccountParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Limit:         10,
		Offset:        0,
	}

	transfers, err := testQueries.GetTransferToFromAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	for _, tf := range transfers {
		valid := (tf.FromAccountID == fromAccount.ID && tf.ToAccountID == toAccount.ID) ||
			(tf.FromAccountID == toAccount.ID && tf.ToAccountID == fromAccount.ID)
		require.True(t, valid)
		require.NotZero(t, tf.ID)
		require.NotZero(t, tf.CreatedAt)

	}
}

func TestUpdateTransfer(t *testing.T) {
	//create transfer
	transfer1 := randomCreateTransfer(t)

	arg := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: utils.RandomMoney(),
	}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, arg.Amount, transfer2.Amount)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.WithinDuration(t, transfer1.CreatedAt.Time, transfer2.CreatedAt.Time, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := randomCreateTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, transfer2)
}
