package domain

import "github.com/shopspring/decimal"

type Balance struct {
	Id         int             `json:"int"`
	Amount     decimal.Decimal `json:"amount"`
	IsFallback bool            `json:"is_fallback"`
}
