package models

type Expense struct {
	ID          int     `db:"id"`
	UserID      int     `db:"id"`
	Description string  `db:"description"`
	Amount      float64 `db:"amount"`
	Date        string  `db:"date"`
}
