package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"strings"
	"time"
)

type accountService struct {
	accountRepo repository.AccountRepository // refer to port
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return accountService{accountRepo: accountRepo}
}

func (s accountService) GetAccounts(customerID int) ([]AccountResponse, error) {
	accounts, err := s.accountRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	responses := []AccountResponse{}
	for _, account := range accounts {
		responses = append(responses, AccountResponse{
			AccountID:   account.AccountID,
			AccountType: account.AccountType,
			OpeningDate: account.OpeningDate,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}
	return responses, nil
}

func (s accountService) NewAccount(customerID int, accountReq NewAccountRequest) (*AccountResponse, error) {
	// Validate
	if accountReq.Amount < 5000 {
		return nil, errs.NewValidationError("amount at least 5,000 baht")
	}
	if strings.ToLower(accountReq.AccountType) != "saving" && strings.ToLower(accountReq.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type should be saving or checking")
	}

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-01-2 15:04:05"),
		AccountType: accountReq.AccountType,
		Amount:      accountReq.Amount,
		Status:      1,
	}
	newAccount, err := s.accountRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	response := AccountResponse{
		AccountID:   newAccount.AccountID,
		AccountType: newAccount.AccountType,
		OpeningDate: newAccount.OpeningDate,
		Amount:      newAccount.Amount,
		Status:      newAccount.Status,
	}
	return &response, nil
}
