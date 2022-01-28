package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepository {
	customers := []Customer{
		{CustomerID: 1001, Name: "Joe", DateOfBirth: "1988-05-21", City: "Chiang Mai", Zipcode: "50000", Status: 1},
		{CustomerID: 1002, Name: "Mikel", DateOfBirth: "1988-05-21", City: "Bangkok", Zipcode: "10310", Status: 1},
	}
	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {

	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}
