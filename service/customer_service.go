package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
)

type customerService struct {
	customerRepo repository.CustomerRepository // refer to port
}

func NewCustomerService(customerRepo repository.CustomerRepository) CustomerService {
	return customerService{customerRepo: customerRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {

	customers, err := s.customerRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	customerResponses := []CustomerResponse{}
	for _, customer := range customers {
		customerResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		customerResponses = append(customerResponses, customerResponse)
	}

	return customerResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.customerRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	customerResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &customerResponse, nil
}
