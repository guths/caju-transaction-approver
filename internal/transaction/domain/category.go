package domain

type Category struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	IsFallback bool   `json:"is_fallback"`
}
