package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func (t *testing.T, got Bitcoin, want Bitcoin) {
		if got != want {
			t.Errorf("got %s want %s", got, want)
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
}
