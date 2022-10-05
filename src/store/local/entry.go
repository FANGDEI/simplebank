package local

import "time"

type Entry struct {
	ID        int64     `json:"id"`
	AccountID int64     `json:"account_id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type SimpleEntry struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

var entries = "entries"

func (m *Manager) CreateEntry(entry SimpleEntry) (Entry, error) {
	var new Entry
	err := m.handler.Table(entries).Create(&entry).Error
	if err != nil {
		return new, err
	}

	err = m.handler.Table(entries).Where("account_id = ? and amount = ?", entry.AccountID, entry.Amount).Take(&new).Error
	if err != nil {
		return new, err
	}
	return new, err
}

func (m *Manager) GetEntry(id int64) (Entry, error) {
	var entry Entry
	err := m.handler.Table(entries).Where("id = ?", id).Take(&entry).Error
	return entry, err
}

func (m *Manager) ListEntries(accountId int64, limit, offset int) ([]Entry, error) {
	list := make([]Entry, 0)
	err := m.handler.Table(entries).Where("account_id = ?", accountId).Limit(limit).Offset(offset).Order("id").Find(&list).Error
	return list, err
}
