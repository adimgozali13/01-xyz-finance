package transaction

import (
	"01-xyz-finance/internal/customerlimit"

)

type Service interface {
	GetAllTransactions() ([]Transaction, error)
	GetTransactionByID(id uint) (*Transaction, error)
	GetLimitTerm(term int, customerID uint) (*customerlimit.CustomerLimit, error)
	UpdateTermLimitAmount(term int, customerID uint, amount float64) (*customerlimit.CustomerLimit, error)
	PaidTransactionAmount(ContractNumber string) (*Transaction, *customerlimit.CustomerLimit, error)
	CreateTransaction(transaction *Transaction) error
	UpdateTransaction(transaction *Transaction) error
	DeleteTransaction(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAllTransactions() ([]Transaction, error) {
	return s.repo.FindAll()
}

func (s *service) GetTransactionByID(id uint) (*Transaction, error) {
	return s.repo.FindByID(id)
}
func (s *service) GetLimitTerm(term int, customerID uint) (*customerlimit.CustomerLimit, error) {
	return s.repo.FindTermLimit(term, customerID)
}

func (s *service) UpdateTermLimitAmount(term int, customerID uint, amount float64) (*customerlimit.CustomerLimit, error) {
	return s.repo.UpdateTermLimit(term, customerID, amount)
}

func (s *service) PaidTransactionAmount(ContractNumber string) (*Transaction, *customerlimit.CustomerLimit, error) {
	return s.repo.PaidTransaction(ContractNumber)
}

func (s *service) CreateTransaction(transaction *Transaction) error {
	// Business logic (e.g., validation) can be added here
	return s.repo.Create(transaction)
}

func (s *service) UpdateTransaction(transaction *Transaction) error {
	// Business logic (e.g., validation) can be added here
	return s.repo.Update(transaction)
}

func (s *service) DeleteTransaction(id uint) error {
	return s.repo.Delete(id)
}
