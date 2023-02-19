package repository

import (
	"example_onion/domain/entity"
	"example_onion/infrastructure/persistence"
)

type UserRepository struct {
	Persistence persistence.UserPersistence
}

func (repo *UserRepository) FindByID(id int) (*entity.User, error) {
	user, err := repo.Persistence.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserRepository(persistence persistence.UserPersistence) *UserRepository {
	return &UserRepository{Persistence: persistence}
}
