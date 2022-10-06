package local

import (
	"fmt"

	"gorm.io/gorm"
)

type Store struct {
	handler *gorm.DB
}

func NewStore() (*Store, error) {
	m, err := New()
	if err != nil {
		return nil, err
	}
	return &Store{
		handler: m.handler,
	}, err
}

// execTx use the function to execute transaction
func (s *Store) execTx(fn func(*Manager) error) error {
	tx := s.handler.Begin()

	m, err := New()
	if err != nil {
		return err
	}

	err = fn(m)
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

		return nil
	})

	return result, err
}
