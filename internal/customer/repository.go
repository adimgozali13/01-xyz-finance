package customer

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Customer, error)
	FindAllWithLimitCust() ([]Customer, error)
	FindByID(id uint) (*Customer, error)
	FindByNik(nik string) (*Customer, error)
	Create(customer *Customer) error
	Update(customer *Customer) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Customer, error) {
	var customers []Customer
	err := r.db.Preload("CustomerLimit").Find(&customers).Error
	return customers, err
}
func (r *repository) FindAllWithLimitCust() ([]Customer, error) {
	var customers []Customer
	err := r.db.Preload("CustomerLimit").Find(&customers).Error
	return customers, err
}

func (r *repository) FindByID(id uint) (*Customer, error) {
	var customer Customer
	err := r.db.First(&customer, id).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
func (r *repository) FindByNik(nik string) (*Customer, error) {
	var customer Customer
	err := r.db.Where("nik = ?", nik).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *repository) Create(customer *Customer) error {
	return r.db.Create(customer).Error
}

func (r *repository) Update(customer *Customer) error {
	return r.db.Save(customer).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Customer{}, id).Error
}
