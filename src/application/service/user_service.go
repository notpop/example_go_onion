package service

import (
	"example_onion/domain/entity"
	"example_onion/domain/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) FindUserByID(id int) (*entity.User, error) {
	return s.UserRepository.FindByID(id)
}
