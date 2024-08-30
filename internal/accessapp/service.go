package accessapp

type Service interface {
	GetAllApps() ([]AccessApp, error)
	// Other service methods
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAllApps() ([]AccessApp, error) {
	return s.repo.FindAll()
}
