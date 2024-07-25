package domain

type BalanceRepository interface {
	GetBalanceByAccountId(categoryId int, accountId int) (Balance, error)
}

type MccRepository interface {
	GetCategoryByMcc(mcc string) (Category, error)
}
