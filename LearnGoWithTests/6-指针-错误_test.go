package main

import (
	"errors"
	"fmt"
	"testing"
)

type Bitcoin int

// 钱包
type Wallet struct {
	//余额
	bitcoin Bitcoin
}

// 放入
func (w *Wallet) Deposit(i Bitcoin) {
	w.bitcoin = w.bitcoin + i
}

// 余额
func (w *Wallet) Balance() Bitcoin {
	return w.bitcoin
}

func (w *Wallet) Withdraw(bitcoin Bitcoin) error {
	if bitcoin > w.bitcoin {
		return errors.New("oh no")
	}

	w.bitcoin -= bitcoin
	return nil
}
func Test_Point_Panic(t *testing.T) {
	t.Run("存入Wallet", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("从Wallet取出", func(t *testing.T) {
		wallet := Wallet{}

		err := wallet.Withdraw(Bitcoin(10))
		if err != nil {
			return
		}

		got := wallet.Balance()
		want := Bitcoin(-10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("", func(t *testing.T) {
		t.Run("Deposit", func(t *testing.T) {
			wallet := Wallet{}
			wallet.Deposit(Bitcoin(10))
			assertBalance(t, wallet, Bitcoin(10))
		})

		t.Run("Withdraw", func(t *testing.T) {
			wallet := Wallet{bitcoin: Bitcoin(20)}
			err := wallet.Withdraw(Bitcoin(10))
			assertBalance(t, wallet, Bitcoin(10))
			assertNoError(t, err)
		})
		var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")
		t.Run("Withdraw insufficient funds", func(t *testing.T) {
			wallet := Wallet{Bitcoin(20)}
			err := wallet.Withdraw(Bitcoin(100))
			assertBalance(t, wallet, Bitcoin(20))
			assertError(t, err, InsufficientFundsError)
		})
	})

}
func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didnt want one")
	}
}
func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
