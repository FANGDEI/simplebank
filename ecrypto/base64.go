package ecrypto

import (
	"encoding/base64"
	"encoding/json"
)

func (m *Manager) MarshlBase64(v any) (string, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf), nil
}

func (m *Manager) UnmarshlBase64(str string, v any) error {
	buf, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, v)
}
