package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	checkBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		if got := wallet.Balance(); got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	}
	assertErrors := func(t *testing.T, error error, want string) {
		t.Helper()
		if error == nil {
			t.Fatal("should fail")
		}

		if error.Error() != want {
			t.Errorf("got %q, but want %q", error.Error(), want)
		}

	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		checkBalance(t, wallet, want)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)

		checkBalance(t, wallet, want)
	})

	t.Run("withdraw with insufficent funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(2)}

		err := wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(2)

		assertErrors(t, err, "insufficent funds")
		checkBalance(t, wallet, want)
	})
}
