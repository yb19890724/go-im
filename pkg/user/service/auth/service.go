package auth

// create repository interface
type Repository interface {
	User(params User) (User, error)
	Add(params User) (int, error)
	Update(params User) (int, error)
}

// create service interface
type Service interface {
	Login(params User) (User, error)
	Register(params User) (int, error)
}

// struct service storage repository
type service struct {
	GR Repository // 寄存器
}

// new service
func NewService(r Repository) Service {
	return &service{r}
}

// service call repository (user and update)
func (s *service) Login(params User) (User, error) {
	return s.GR.User(params)
}

// service call repository Add
func (s *service) Register(params User) (int, error) {
	return s.GR.Add(params)
}
