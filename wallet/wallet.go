package wallet

import "fmt"

type Wallet struct {
	balance int
}

func (d *Wallet) Deposit(amount int) {
	if amount <= 0 {
		return
	}
	d.balance += amount

}

func (w *Wallet) Withdraw(amount int) error {
	if amount > w.balance {
		return fmt.Errorf("Your amount is very high")
	}
	w.balance -= amount
	return nil
}

func (b *Wallet) Balance() int {
	return b.balance
}
