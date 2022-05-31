package event

type Service interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	Del(id int64) error
	Update(id int64, name string) error
	Create(event *Event) error
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]Event, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindOne(id int64) (*Event, error) {
	return (*s.repo).FindOne(id)
}

func (s *service) Del(id int64) error{
	return (*s.repo).Del(id)
}

func (s *service) Update(id int64, name string) error{
	return (*s.repo).Update(id, name)
}

func (s *service) Create(event *Event) error{
	return (*s.repo).Create(event)
}