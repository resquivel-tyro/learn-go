package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func (t *testing.T, got Bitcoin, want Bitcoin) {
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertError := func (t testing.TB, err error) {
		// nil is synonymous with null from other programming languages.
		// Errors can be nil because the return type of Withdraw will be error, which is an interface.
		// If you see a function that takes arguments or returns values that are interfaces, they can be nillable.
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet {}

		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)
		assertBalance(t, got, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet {balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)
		assertBalance(t, got, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet {balance: startingBalance}

		err := wallet.Withdraw(Bitcoin(21))

		got := wallet.Balance()
		assertError(t, err)
		assertBalance(t, got, startingBalance)
	})
}
