package pointer_errors

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {

	assertNoError := func(t *testing.T, got error) {

	}
	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if !errors.Is(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.WithDraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalnace := Bitcoin(20)
		wallet := Wallet{startingBalnace}
		err := wallet.WithDraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalnace)
		assertError(t, err, ErrInsufficientFunds)
	})
}
