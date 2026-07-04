package wallet

import "testing"

func TestDepositFile(t *testing.T) {
	t.Run("Два пополнения", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(100)
		w.Deposit(50)
		got := w.Balance()
		want := 150
		if got != want {
			t.Errorf("получили %d, ожидали %d", got, want)
		}
	})
	t.Run("Одно списание", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(100)
		w.Withdraw(30)
		got := w.Balance()
		want := 70
		if got != want {
			t.Errorf("получили %d, ожидали %d", got, want)
		}
	})
	t.Run("Проверка ошибки", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(100)
		err := w.Withdraw(200)
		if err == nil {
			t.Errorf("ожидали ошибку, но получили nil")
		}
	})
	t.Run("Отрицательный депозит не меняет баланс", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(-100)
		got := w.Balance()
		if got != 0 {
			t.Errorf("получили %d, ожидали %d", got, 0)
		}
	})
}
