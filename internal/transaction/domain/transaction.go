package domain

type Type string

var (
	Debit  Type = "debit"
	Credit Type = "credit"
)

type Transaction struct {
	Id     int  `json:"id"`
	Amount int  `json:"amount"`
	Type   Type `json:"type"`
}
