package domain

type Wallet struct {
	ID      int
	UserID  int
	Balance float64
}

func (w Wallet) HasBalance(amount float64) bool {
	return w.Balance >= amount
}

func (w *Wallet) SetId(id int) {
	if w.ID == 0 {
		w.ID = id
	}
}
