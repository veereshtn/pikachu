package svc

import (
	db_commons "github.com/kutty-kumar/db_commons/model"
	"pikachu/pkg/repository"
)

type UserAttributeService struct {
	db_commons.BaseSvc
	UserAttributeService repository.UserAttributeRepository
}

func NewUserAttributeService(baseSvc db_commons.BaseSvc, userAttributeRepository repository.UserAttributeRepository) UserAttributeService {
	return UserAttributeService{
		baseSvc,
		userAttributeRepository,
	}
}


