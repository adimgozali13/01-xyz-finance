package accessapp

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]AccessApp, error)
	FindByID(id uint) (*AccessApp, error)
	Create(app *AccessApp) error
	Update(app *AccessApp) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]AccessApp, error) {
	var apps []AccessApp
	err := r.db.Find(&apps).Error
	return apps, err
}

// FindByID retrieves a single AccessApp record by ID
func (r *repository) FindByID(id uint) (*AccessApp, error) {
	var accessApp AccessApp
	err := r.db.First(&accessApp, id).Error
	if err != nil {
		return nil, err
	}
	return &accessApp, nil
}

// Create inserts a new AccessApp record into the database
func (r *repository) Create(accessApp *AccessApp) error {
	return r.db.Create(accessApp).Error
}

// Update modifies an existing AccessApp record
func (r *repository) Update(accessApp *AccessApp) error {
	return r.db.Save(accessApp).Error
}

// Delete removes an AccessApp record by ID
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&AccessApp{}, id).Error
}
