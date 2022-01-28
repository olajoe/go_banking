package service

type AccountResponse struct {
	AccountID   int     `json:"account_id"`
	AccountType string  `json:"account_type"`
	OpeningDate string  `json:"opening_date"`
	Amount      float64 `json:"amount"`
	Status      int     `json:"status"`
}

type NewAccountRequest struct {
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

type AccountService interface {
	GetAccounts(int) ([]AccountResponse, error)
	NewAccount(int, NewAccountRequest) (*AccountResponse, error)
}
