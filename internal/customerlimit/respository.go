package customerlimit

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]CustomerLimit, error)
	FindByID(id uint) (*CustomerLimit, error)
	FindByTerm(term int, customerID uint) (*CustomerLimit, error)
	Create(customerLimit *CustomerLimit) error
	Update(customerLimit *CustomerLimit) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]CustomerLimit, error) {
	var customerLimits []CustomerLimit
	err := r.db.Find(&customerLimits).Error
	return customerLimits, err
}

func (r *repository) FindByID(id uint) (*CustomerLimit, error) {
	var customerLimit CustomerLimit
	err := r.db.First(&customerLimit, id).Error
	if err != nil {
		return nil, err
	}
	return &customerLimit, nil
}
func (r *repository) FindByTerm(term int, customerID uint ) (*CustomerLimit, error) {
	var customerLimit CustomerLimit
	err := r.db.Where("term = ? AND customer_id = ?", term, customerID).First(&customerLimit).Error
	if err != nil {
		return nil, err
	}
	return &customerLimit, nil
}

func (r *repository) Create(customerLimit *CustomerLimit) error {
	return r.db.Create(customerLimit).Error
}

func (r *repository) Update(customerLimit *CustomerLimit) error {
	return r.db.Save(customerLimit).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&CustomerLimit{}, id).Error
}
