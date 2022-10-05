package local

import "time"

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

func (m *Manager) CreateAccount(account SimpleAccount) (Account, error) {
	var new Account
	err := m.handler.Table(accounts).Create(&account).Error
	if err != nil {
		return new, err
	}

	err = m.handler.Table(accounts).Where("owner = ? and balance = ? and currency = ?", account.Owner, account.Balance, account.Currency).Take(&new).Error
	if err != nil {
		return new, err
	}
	return new, err
}

func (m *Manager) GetAccount(id int64) (Account, error) {
	var account Account
	err := m.handler.Table(accounts).Where("id = ?", id).Take(&account).Error
	return account, err
}

func (m *Manager) ListAccounts(limit, offset int) ([]Account, error) {
	list := make([]Account, 0)
	err := m.handler.Table(accounts).Limit(limit).Offset(offset).Order("id").Find(&list).Error
	return list, err
}

func (m *Manager) UpdateAccount(id, balance int64) (Account, error) {
	var new Account
	err := m.handler.Table(accounts).Where("id = ?", id).Update("balance", balance).Error
	if err != nil {
		return new, err
	}

	err = m.handler.Table(accounts).Where("id = ?", id).Take(&new).Error
	if err != nil {
		return new, err
	}
	return new, nil
}

func (m *Manager) DeleteAccount(id int64) error {
	return m.handler.Table(accounts).Where("id = ?", id).Delete(&Account{}).Error
}
