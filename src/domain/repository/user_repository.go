package repository

import (
	"example_onion/domain/entity"
	"example_onion/infrastructure/persistence"
)

type UserRepository struct {
	persistence persistence.UserPersistence
}

func (repo *UserRepository) FindByID(id int) (*entity.User, error) {
	user, err := repo.persistence.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserRepository(persistence persistence.UserPersistence) UserRepository {
	return UserRepository{persistence: persistence}
}
