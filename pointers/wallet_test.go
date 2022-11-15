package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Tests you can deposit into the wallet", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		expect := Bitcoin(10)

		assertBalance(t, wallet, expect)
	})

	t.Run("Tests you can withdraw from the wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(5))

		expect := Bitcoin(15)

		assertNoError(t, err)
		assertBalance(t, wallet, expect)
	})

	t.Run("Tests withdrawind insufficient funds", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{initialBalance}

		err := wallet.Withdraw(Bitcoin(50))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, initialBalance)

	})
}

func assertBalance(t testing.TB, wallet Wallet, expect Bitcoin) {
	t.Helper()

	result := wallet.Balance()

	if result != expect {
		t.Errorf("\nresult: %s\nexpect: %s\n", result, expect)
	}

}

func assertError(t testing.TB, returned, expected error) {
	t.Helper()

	// Using Fatal as we do not want to make more assertions on the error returned if there isn't an error. Without this the test would carry on and panic because of a nil pointer.
	if returned == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if returned != expected {
		t.Errorf("result: %q\nexpect: %q", returned, expected)
	}
}

func assertNoError(t testing.TB, returned error) {
	t.Helper()

	if returned != nil {
		t.Fatal("got an error but didn't want one")
	}
}
