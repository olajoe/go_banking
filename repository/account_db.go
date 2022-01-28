package repository

import "github.com/jmoiron/sqlx"

type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository {
	return accountRepositoryDB{db: db}
}

func (r accountRepositoryDB) Create(accInfo Account) (*Account, error) {
	query := "INSERT INTO accounts (customer_id,opening_date,account_type,amount,status) VALUES(?,?,?,?,?)"
	result, err := r.db.Exec(
		query,
		accInfo.CustomerID,
		accInfo.OpeningDate,
		accInfo.AccountType,
		accInfo.Amount,
		accInfo.Status,
	)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	accInfo.AccountID = int(id)
	return &accInfo, nil
}

func (r accountRepositoryDB) GetAll(customerID int) ([]Account, error) {
	accounts := []Account{}
	query := "SELECT account_id,customer_id, opening_date, account_type, amount, status FROM accounts WHERE customer_id=?"
	err := r.db.Select(&accounts, query, customerID)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
