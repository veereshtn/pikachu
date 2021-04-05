package repository

import "pikachu/pkg/domain"

type UserRepository interface {
	Create(user *domain.User) (error, *domain.User)
	Update(id string, user *domain.User) (error, *domain.User)
	FindByExternalId(id string)(error, *domain.User)
	MultiGetByExternalIds(ids []string)(error, []domain.User)
}
