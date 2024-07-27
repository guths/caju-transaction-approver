package domain

import "github.com/shopspring/decimal"

type Balance struct {
	Id         int             `json:"int"`
	Amount     decimal.Decimal `json:"amount"`
	AccountId  int             `json:"account_id"`
	CategoryId int             `json:"category_id"`
}
