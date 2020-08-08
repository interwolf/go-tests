package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin is a type of int
type Bitcoin int

// ErrWithdrawTooMuch when withdraw > balance
var ErrWithdrawTooMuch = errors.New("withdraw amount > balance")

// Wallet for bitcoin
type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit adds money to balance
func (w *Wallet) Deposit(deposit Bitcoin) {
	println(w)
	w.balance += deposit
	return
}

// Withdraw reduces money from balance
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrWithdrawTooMuch
	}

	w.balance -= amount
	return nil
}

// Balance shows money of account
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Stringer is an interface of String() func
// type Stringer interface {
// 	String() string
// }
