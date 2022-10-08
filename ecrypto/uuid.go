package ecrypto

import (
	"strings"

	"github.com/google/uuid"
)

func (m *Manager) GetUUID() string {
	return uuid.NewString()
}

func (m *Manager) GetUUIDWithoutSplit() string {
	uuid := m.GetUUID()
	return strings.ReplaceAll(uuid, "-", "")
}
