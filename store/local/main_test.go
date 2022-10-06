package local

import "testing"

var testManager *Manager

func TestMain(m *testing.M) {
	testManager, _ = New()
}
