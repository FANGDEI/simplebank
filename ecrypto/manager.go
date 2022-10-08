package ecrypto

type Manager struct {
	secret	[]byte
}

func New() *Manager {
	return &Manager{
		secret: []byte(C.Secret),
	}
}
