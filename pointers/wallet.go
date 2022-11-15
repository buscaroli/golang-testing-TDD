package pointers

import (
	"errors"
	"fmt"
)

// Creating an Error that can be used when funds are insufficient. Prevents errors due to changing the string.
var ErrInsufficientFunds = errors.New("insufficient funds, unable to withdraw")

type Bitcoin int

// Prints the bitcoin value as a string as '<num> BTC' eg '12 BTC'.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// Deposit bitcoin into the wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Retrieve the Balance of bitcoins kept in the wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw a set amount of bitcoins from the wallet.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
