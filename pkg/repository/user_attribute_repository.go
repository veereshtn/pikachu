package repository

import "pikachu/pkg/domain"

type UserAttributeRepository interface {
	CreateUserAttribute(userId string, userAttribute *domain.UserAttribute) (error, *domain.UserAttribute)
	UpdateUserAttribute(userId string, userAttribute *domain.UserAttribute) (error, *domain.UserAttribute)
	ListUserAttributes(userId string) (error, []domain.UserAttribute)
}