package file

type Service interface {
	Save(file *File) error
	ListByUser(userID uint) ([]File, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Save(file *File) error {
	return s.repo.Save(file)
}

func (s *service) ListByUser(userID uint) ([]File, error) {
	return s.repo.ListByUser(userID)
}
