package local

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Something wrong with the test. Just left it here
func TestCreateAccount(t *testing.T) {
	arg := Account{
		Owner:    RandomOwner(),
		Balance:  RandomMoney(),
		Currency: RandomCurrency(),
	}

	account, err := testManager.CreateAccount(arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
