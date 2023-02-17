package usecase

import (
	"errors"

	"github.com/clean-architecture/entity"
	"github.com/clean-architecture/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(name string, email string) (*entity.User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}

	return s.userRepository.CreateUser(name, email)
}

func (s *UserService) GetUser(id int) (*entity.User, error) {
	return s.userRepository.GetUser(id)
}

func (s *UserService) ListUsers() ([]*entity.User, error) {
	return s.userRepository.ListUsers()
}

// type UserService interface {
// 	CreateUser(name string, email string) (entity.User, error)
// 	GetUser(id int) (*entity.User, error)
// 	ListUsers() ([]*entity.User, error)
// }
