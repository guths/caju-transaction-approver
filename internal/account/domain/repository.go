package domain

type AccountRepository interface {
	GetAccountByIdentifier(accountIdentifier string) (*Account, error)
}
