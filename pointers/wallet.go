package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin represents the number of Bitcoins
type Bitcoin int

// String function implementation that determines how Bitcoin will printed when %s format verb is used
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet stores the number of Bitcoin that someone owes
type Wallet struct {
	balance Bitcoin
}

// Deposit will add `amount` BTC to a wallet
func (w *Wallet) Deposit(amount Bitcoin)  {
	w.balance += amount
}

// ErrInsufficientFunds means wallet does not have sufficient balance to perform a withdrawal
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")


// Withdraw will deduct `amount` BTC to a wallet, returning error if insufficient balance
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance returns the amount of BTC in a wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
