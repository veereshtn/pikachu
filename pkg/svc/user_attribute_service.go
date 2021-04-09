package svc

import (
	db_commons "github.com/kutty-kumar/db_commons/model"
	"pikachu/pkg/domain"
	"pikachu/pkg/repository"
)

type UserAttributeService struct {
	db_commons.BaseSvc
	UserAttributeRepo repository.UserAttributeRepository
}

func NewUserAttributeService(baseSvc db_commons.BaseSvc, userAttributeRepository repository.UserAttributeRepository) UserAttributeService {
	return UserAttributeService{
		baseSvc,
		userAttributeRepository,
	}
}

func (uas *UserAttributeService) UpdateUserAttribute(userId string, other *domain.UserAttribute) (error, *domain.UserAttribute){
	return uas.UserAttributeRepo.UpdateUserAttribute(userId, other)
}

func (uas *UserAttributeService) GetUserAttributesByKey(userId string, attributeKey string) (error, *domain.UserAttribute) {
	return uas.UserAttributeRepo.GetUserAttributeByKey(userId, attributeKey)
}

func (uas *UserAttributeService) ListUserAttributes(userId string)(error, []domain.UserAttribute){
	return uas.UserAttributeRepo.ListUserAttributes(userId)
}


