package customer

type Service interface {
	GetAllCustomers() ([]Customer, error)
	GetAllCustomersWithLimitCust() ([]Customer, error)
	GetCustomerByID(id uint) (*Customer, error)
	GetCustomerByNik(nik string) (*Customer, error)
	CreateCustomer(customer *Customer) error
	UpdateCustomer(customer *Customer) error
	DeleteCustomer(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAllCustomers() ([]Customer, error) {
	return s.repo.FindAll()
}

func (s *service) GetAllCustomersWithLimitCust() ([]Customer, error) {
	return s.repo.FindAllWithLimitCust()
}

func (s *service) GetCustomerByID(id uint) (*Customer, error) {
	return s.repo.FindByID(id)
}
func (s *service) GetCustomerByNik(nik string) (*Customer, error) {
	return s.repo.FindByNik(nik)
}

func (s *service) CreateCustomer(customer *Customer) error {
	return s.repo.Create(customer)
}

func (s *service) UpdateCustomer(customer *Customer) error {
	return s.repo.Update(customer)
}

func (s *service) DeleteCustomer(id uint) error {
	return s.repo.Delete(id)
}
