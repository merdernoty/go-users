package service

import "github.com/merdernoty/microservices-planner/users/internal/user/domain"

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]domain.User, error) {
	return s.repo.GetAll()
}

func (s *userService) GetUser(id int64) (*domain.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) CreateUser(user *domain.User) error {
	return s.repo.Create(user)
}

func (s *userService) UpdateUser(user *domain.User) error {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id int64) error {
	return s.repo.Delete(id)
}

func (s *userService) GetUserByEmail(email string) (*domain.User, error) {
	return s.repo.GetUserByEmail(email)
}