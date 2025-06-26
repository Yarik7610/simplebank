package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (s *Store) execTx(ctx context.Context, f func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil
	}

	q := New(tx)
	err = f(q)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %v", err, rollbackErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (s *Store) TransferTx(ctx context.Context, arg CreateTransferParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := s.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, arg)
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// make always same update order by smallest id
		firstAddAccountBalanceParams, secondAddAccountBalanceParams := getAddAccountBalancesParams(arg.FromAccountID, arg.ToAccountID, arg.Amount)

		result.FromAccount, err = q.AddAccountBalance(context.Background(), firstAddAccountBalanceParams)
		if err != nil {
			return err
		}

		result.ToAccount, err = q.AddAccountBalance(context.Background(), secondAddAccountBalanceParams)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

func getAddAccountBalancesParams(fromAccountID, toAccountID, amount int64) (firstAddAccountBalanceParams, secondAddAccountBalanceParams AddAccountBalanceParams) {
	if fromAccountID < toAccountID {
		firstAddAccountBalanceParams = AddAccountBalanceParams{fromAccountID, -amount}
		secondAddAccountBalanceParams = AddAccountBalanceParams{toAccountID, amount}
	} else {
		firstAddAccountBalanceParams = AddAccountBalanceParams{toAccountID, amount}
		secondAddAccountBalanceParams = AddAccountBalanceParams{fromAccountID, -amount}
	}
	return
}
