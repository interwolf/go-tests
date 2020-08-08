package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		checkBalance(t, &wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(2))
		assertNoError(t, err)

		checkBalance(t, &wallet, Bitcoin(8))
	})

	t.Run("Withdraw too much", func(t *testing.T) {
		initBalance := Bitcoin(10)

		wallet := Wallet{balance: initBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrWithdrawTooMuch)

		checkBalance(t, &wallet, Bitcoin(10))

	})
}

func checkBalance(t *testing.T, wallet *Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("should return an error but not")
	}

	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("should return no error but got %q", got)
	}
}
