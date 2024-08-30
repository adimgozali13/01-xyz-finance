package transaction

import (
	"fmt"
	"gorm.io/gorm"
	"01-xyz-finance/internal/customerlimit"

)

type Repository interface {
	FindAll() ([]Transaction, error)
	FindByID(id uint) (*Transaction, error)
	FindTermLimit(term int, customerID uint) (*customerlimit.CustomerLimit, error)
	UpdateTermLimit(term int, customerID uint, amount float64) (*customerlimit.CustomerLimit, error)
	PaidTransaction(ContractNumber string) (*Transaction, *customerlimit.CustomerLimit, error)
	Create(transaction *Transaction) error
	Update(transaction *Transaction) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Preload("Customer").Preload("CustomerLimit").Find(&transactions).Error
	return transactions, err
}

func (r *repository) FindByID(id uint) (*Transaction, error) {
	var transaction Transaction
	err := r.db.Preload("Customer").Preload("CustomerLimit").First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *repository) FindTermLimit(term int, customerID uint) (*customerlimit.CustomerLimit, error) {

	var cLimit customerlimit.CustomerLimit
	err := r.db.Where("term = ? AND customer_id = ?", term, customerID).First(&cLimit).Error
	
	if err != nil {
		return nil, err
	}
	return &cLimit, nil
}

func (r *repository) UpdateTermLimit(term int, customerID uint, amount float64) (*customerlimit.CustomerLimit, error) {

	var cLimit customerlimit.CustomerLimit
	err := r.db.Model(&cLimit).Where("term = ? AND customer_id = ?", term, customerID).Update("amount", amount).Error
	
	if err != nil {
		return nil, err
	}
	return &cLimit, nil
}

func (r *repository) PaidTransaction(ContractNumber string) (*Transaction, *customerlimit.CustomerLimit, error) {
	var transaction Transaction
	var climit customerlimit.CustomerLimit

	err := r.db.Where("contract_number = ?", ContractNumber).First(&transaction).Error
	if err != nil {
		return nil, nil, err
	}

	if transaction.Status == "Paid" {
		return nil, nil, fmt.Errorf("Customer has already paid")
	}

	err = r.db.Model(&transaction).
		Update("status", "Paid").
		Error

	if err != nil {
		return nil, nil, err
	}

	err = r.db.Model(&climit).
		Where("id = ?", transaction.CustomerLimitID).
		Update("amount", gorm.Expr("amount + ?", transaction.OTR)).
		Error

	if err != nil {
		return nil, nil, err
	}

	return &transaction, &climit, nil
}



func (r *repository) Create(transaction *Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *repository) Update(transaction *Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Transaction{}, id).Error
}
