package local

import (
	"fmt"

	"gorm.io/gorm"
)

type Store struct {
	handler *gorm.DB
}

func newStore(handler *gorm.DB) *Store {
	return &Store{
		handler: handler,
	}
}

// execTx use the function to execute transaction
func (s *Store) execTx(fn func(*Manager) error) error {
	tx := s.handler.Begin()

	m := new(tx)
	err := fn(m)
	if err != nil {
		if rbErr := tx.Rollback().Error; rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit().Error
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (s *Store) TransferTx(arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := s.execTx(func(m *Manager) error {
		var err error

		result.Transfer, err = m.CreateTransfer(Transfer{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return nil
		}

		result.FromEntry, err = m.CreateEntry(Entry{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = m.CreateEntry(Entry{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// TODO: update accounts' balance
		account1, err := m.GetAccountForUpdate(arg.FromAccountID)
		if err != nil {
			return err
		}

		result.FromAccount, err = m.UpdateAccount(Account{
			ID:        arg.FromAccountID,
			Owner:     account1.Owner,
			Balance:   account1.Balance - arg.Amount,
			Currency:  account1.Currency,
			CreatedAt: account1.CreatedAt,
		})
		if err != nil {
			return err
		}

		account2, err := m.GetAccountForUpdate(arg.ToAccountID)
		if err != nil {
			return err
		}

		result.ToAccount, err = m.UpdateAccount(Account{
			ID:        arg.ToAccountID,
			Owner:     account2.Owner,
			Balance:   account2.Balance + arg.Amount,
			Currency:  account2.Currency,
			CreatedAt: account2.CreatedAt,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
