package domain

import "github.com/shopspring/decimal"

type BalanceRepository interface {
	DebitAmount(accountId int, categoryId int, amount decimal.Decimal) (*Transaction, error)
	GetBalance(accountId int, categoryId int) (*Balance, error)
}

type MccRepository interface {
	GetCategoryByMcc(mcc string) (Category, error)
	GetFallbackCategory() (*Category, error)
}
