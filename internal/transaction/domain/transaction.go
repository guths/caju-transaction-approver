package domain

import "github.com/shopspring/decimal"

type Type string

var (
	Debit  Type = "debit"
	Credit Type = "credit"
)

type Transaction struct {
	Id     int             `json:"id"`
	Amount decimal.Decimal `json:"amount"`
	Type   Type            `json:"type"`
}
