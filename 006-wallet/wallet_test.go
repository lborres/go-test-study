package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, expected Bitcoin) {
		t.Helper()

		balance := wallet.Balance()
		if balance != expected {
			t.Errorf("expected %s got %s", expected, balance)
		}
	}

	assertError := func(t testing.TB, err, expected error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
		if err != expected {
			t.Errorf("expected %q got %q", expected, err)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(45694)

		expected := Bitcoin(45694)

		assertBalance(t, wallet, expected)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(45694)}
		wallet.Withdraw(10000)

		expected := Bitcoin(35694)

		assertBalance(t, wallet, expected)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(45694)
		wallet := Wallet{balance: startBalance}
		err := wallet.Withdraw(50000)

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startBalance)
	})
}
