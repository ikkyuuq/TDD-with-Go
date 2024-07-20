package pointererror

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(n Bitcoin) {
	w.balance += n
}

var ErrInsufficientFunds = errors.New("not enough balance to make transaction!")

func (w *Wallet) Withdraw(n Bitcoin) error {
	if n > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= n
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
