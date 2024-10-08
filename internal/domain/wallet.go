package domain

type Wallet struct {
	ID      int
	UserID  int
	Balance float64
}

func (w *Wallet) HasBalance(amount float64) bool {
	return w.Balance >= amount
}

func (w *Wallet) SetId(id int) {
	if w.ID == 0 {
		w.ID = id
	}
}

func (w *Wallet) RemoveBalance(amount float64) {
	w.Balance = w.Balance - amount
}

func (w *Wallet) AddBalance(amount float64) {
	w.Balance = w.Balance + amount
}
