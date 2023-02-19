package usecase

import (
    "example_onion/domain/entity"
)

type UserUsecase interface {
    GetUserInfo(id int) (*entity.User, error)
}
