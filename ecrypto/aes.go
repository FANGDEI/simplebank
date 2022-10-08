package ecrypto

import (
	"encoding/json"

	"github.com/wumansgy/goEncrypt/aes"
)

func (m *Manager) EncryptAesJson(v any) (string, error) {
	return m.EncryptAesJsonWithSecret(v, m.secret)
}

func (m *Manager) DecryptAesJson(hex string, v any) error {
	return m.DecryptAesJsonWithSecret(hex, v, m.secret)
}

func (m *Manager) EncryptAesString(str string) (string, error) {
	return m.EncryptAesStringWithSecret(str, m.secret)
}

func (m *Manager) DecryptAesString(hex string) (string, error) {
	return m.DecryptAesStringWithSecret(hex, m.secret)
}

func (m *Manager) GetSecret(appID string) []byte {
	size := len(appID)
	if size < 16 {
		for i := 0; i < 16 - size; i ++ {
			appID += "x"
		}
		return []byte(appID)
	}
	return []byte(appID[:16])
}

func (m *Manager) EncryptAesJsonWithSecret(v any, secret []byte) (string, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return aes.AesCbcEncryptHex(buf, secret, m.secret)
}

func (m *Manager) DecryptAesJsonWithSecret(hex string, v any, secret []byte) error {
	buf, err := aes.AesCbcDecryptByHex(hex, secret, m.secret)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, v)
}

func (m *Manager) EncryptAesStringWithSecret(str string, secret []byte) (string, error) {
	return aes.AesCbcEncryptHex([]byte(str), secret, m.secret)
}

func (m *Manager) DecryptAesStringWithSecret(hex string, secret []byte) (string, error) {
	buf, err := aes.AesCbcDecryptByHex(hex, secret, m.secret)
	return string(buf), err
}
