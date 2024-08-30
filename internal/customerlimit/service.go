package customerlimit

type Service interface {
	GetAllCustomerLimits() ([]CustomerLimit, error)
	GetCustomerLimitByID(id uint) (*CustomerLimit, error)
	GetCustomerLimitByTerm(term int, customerID uint) (*CustomerLimit, error)
	CreateCustomerLimit(customerLimit *CustomerLimit) error
	UpdateCustomerLimit(customerLimit *CustomerLimit) error
	DeleteCustomerLimit(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAllCustomerLimits() ([]CustomerLimit, error) {
	return s.repo.FindAll()
}

func (s *service) GetCustomerLimitByID(id uint) (*CustomerLimit, error) {
	return s.repo.FindByID(id)
}

func (s *service) GetCustomerLimitByTerm(term int, customerID uint) (*CustomerLimit, error) {
	return s.repo.FindByTerm(term, customerID)
}

func (s *service) CreateCustomerLimit(customerLimit *CustomerLimit) error {
	return s.repo.Create(customerLimit)
}

func (s *service) UpdateCustomerLimit(customerLimit *CustomerLimit) error {
	return s.repo.Update(customerLimit)
}

func (s *service) DeleteCustomerLimit(id uint) error {
	return s.repo.Delete(id)
}
