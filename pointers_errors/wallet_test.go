package wallet

import (
	"testing"
)

func Test_Wallet(t *testing.T) {

	AssertEquals := func(t *testing.T, want Bitcoin, got Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("Got %s, but wanted %s", got, want)
		}
	}

	t.Run("Deposit tests", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		AssertEquals(t, want, got)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		wallet.Withdraw(5)

		got := wallet.Balance()
		want := Bitcoin(5)

		AssertEquals(t, want, got)
	})

	t.Run("Cannot withdraw more, than waht you have", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		err := wallet.Withdraw(100)

		if err == nil {
			t.Error(`You must return an error if you try to withdraw more, than\
			what you have!`)
		}
	})
}
