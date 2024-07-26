package domain

type Account struct {
	Id                int    `json:"id"`
	AccountIdentifier string `json:"account_identifier"`
	Name              string `json:"name"`
}
