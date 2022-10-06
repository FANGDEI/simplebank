package local

import (
	"time"

	"gorm.io/gorm/clause"
)

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type SimpleAccount struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

var accounts = "accounts"

func (m *Manager) CreateAccount(account Account) (Account, error) {
	err := m.handler.Table(accounts).Create(&account).Error
	return account, err
}

func (m *Manager) GetAccount(id int64) (Account, error) {
	var account Account
	err := m.handler.Table(accounts).Where("id = ?", id).Take(&account).Error
	return account, err
}

func (m *Manager) GetAccountForUpdate(id int64) (Account, error) {
	var account Account
	err := m.handler.Table(accounts).Clauses(clause.Locking{Strength: "NO KEY UPDATE"}).Where("id = ?", id).Take(&account).Error
	return account, err
}

func (m *Manager) ListAccounts(limit, offset int) ([]Account, error) {
	list := make([]Account, 0)
	err := m.handler.Table(accounts).Limit(limit).Offset(offset).Order("id").Find(&list).Error
	return list, err
}

func (m *Manager) UpdateAccount(account Account) (Account, error) {
	err := m.handler.Table(accounts).Where("id = ?", account.ID).Updates(&account).Error
	return account, err
}

func (m *Manager) DeleteAccount(id int64) error {
	return m.handler.Table(accounts).Where("id = ?", id).Delete(&Account{}).Error
}

func (m *Manager) AddAccountBalance(account Account) (Account, error) {
	return m.UpdateAccount(account)
}
