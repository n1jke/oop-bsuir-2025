package domain

type Money struct {
	amount int
	curr   Currency
}

type Currency string

func NewMoney(amount int, curr Currency) Money {
	return Money{amount: amount, curr: curr}
}

func (m Money) Add(other Money) Money {
	return Money{amount: m.amount + other.amount, curr: m.curr}
}

func (m Money) Sub(other Money) Money {
	return Money{amount: m.amount - other.amount, curr: m.curr}
}
