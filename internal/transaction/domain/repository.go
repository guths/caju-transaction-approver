package domain

type BalanceRepository interface {
}

type MccRepository interface {
	GetCategoryByMcc(mcc string) Category
}
