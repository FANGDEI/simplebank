package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Manager struct {
	handler *redis.Client
}

func New() (*Manager, error) {
	m := &Manager{
		handler: redis.NewClient(
			&redis.Options{
				Addr:     C.Addr,
				Password: C.Password,
			},
		),
	}
	return m, m.handler.Ping(context.Background()).Err()
}

func (m *Manager) getNormalSessionKey(equip string, id int64) string {
	return fmt.Sprintf(
		"session_normal_%d_%s",
		id,
		equip,
	)
}

func (m *Manager) getEmailVerifyKey(account string) string {
	return fmt.Sprintf(
		"email_code_%s",
		account,
	)
}

func (m *Manager) getEmailAccountKey(account string) string {
	return fmt.Sprintf(
		"email_ban_%s",
		account,
	)
}

func (m *Manager) getSSOMobileKey(uuid string) string {
	return fmt.Sprintf(
		"sso_mobile_%s",
		uuid,
	)
}

func (m *Manager) getSSOUserKey(id int64, secretID string) string {
	return fmt.Sprintf(
		"sso_user_%d_%s",
		id,
		secretID,
	)
}

func (m *Manager) getSSOUserAllKey(id int64) string {
	return fmt.Sprintf(
		"sso_user_%d_*",
		id,
	)
}

func (m *Manager) getNormalAllSessionKey(id int64) string {
	return fmt.Sprintf(
		"session_normal_%d_*",
		id,
	)
}

func (m *Manager) getAuthUserKey(code string) string {
	return fmt.Sprintf(
		"auth_user_%s",
		code,
	)
}

func (m *Manager) getClientMobileKey(id int64) string {
	return fmt.Sprintf(
		"client_mobile_%d",
		id,
	)
}
