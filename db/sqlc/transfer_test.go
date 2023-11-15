package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	arg := CreateTransfersParams{

		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        account1.Balance,
	}

	transfer, err := testQueries.CreateTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.Amount, transfer.Amount)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer

}
func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)

}

func TestGetTransfers(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	transfer2, err := testQueries.GetTransfers(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

// func TestUpdateTransfer(t *testing.T){
// 	tranfer1 := createRandomTransfer(t)

// 	arg := UpdateTransferParams{
// 		ID: transfer1.ID,
// 		Amount: util.RandomMoney().
// 	}

// 	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, transfer2)

// 	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
// 	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
// 	require.Equal(t, transfer1.ID, transfer2.ID)
// 	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
// }

func TestDeleteTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)

}
