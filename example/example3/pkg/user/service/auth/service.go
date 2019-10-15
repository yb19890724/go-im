package auth

// create repository interface
type Repository interface {
	User(ps User) (User, error)
}

// create service interface
type Service interface {
	User(ps User) (User, error)
}

// struct service storage repository
type service struct {
	GR Repository // 寄存器
}

// new service
func NewService(r Repository) Service {
	return &service{r}
}

// service call repository (user and update)u
func (s *service) User(ps User) (u User, err error) {
	res, err := s.GR.User(ps)
	return res, err
}
