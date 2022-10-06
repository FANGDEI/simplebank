package local

import "time"

type Transfer struct {
	ID            int64     `json:"id"`
	FromAccountID int64     `json:"from_account_id"`
	ToAccountID   int64     `json:"to_account_id"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

type SimpleTransfer struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

var transfers = "transfers"

func (m *Manager) CreateTransfer(transfer Transfer) (Transfer, error) {
	err := m.handler.Table(transfers).Create(&transfer).Error
	return transfer, err
}

func (m *Manager) GetTransfer(id int64) (Transfer, error) {
	var transfer Transfer
	err := m.handler.Table(transfers).Where("id = ?", id).Take(&transfer).Error
	return transfer, err
}

func (m *Manager) ListTransfers(fromAccountId, toAccountId int64, limit, offset int) ([]Transfer, error) {
	list := make([]Transfer, 0)
	err := m.handler.Table(entries).Where("from_account_id = ? and to_account_id = ?", fromAccountId, toAccountId).Limit(limit).Offset(offset).Order("id").Find(&list).Error
	return list, err
}
