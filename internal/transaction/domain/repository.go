package domain

import "github.com/shopspring/decimal"

type BalanceRepository interface {
	DebitAmount(accountId int, categoryId int, amount decimal.Decimal) (*Transaction, error)
}

type MccRepository interface {
	GetCategoryByMcc(mcc string) (Category, error)
	GetFallbackCategory() (*Category, error)
}
