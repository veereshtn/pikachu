package repository

import (
	"pikachu/pkg/domain"
)

type AddressRepository interface {
	CreateUserAddress(userId string, address *domain.Address) (error, *domain.Address)
	UpdateUserAddress(userId string, addressId string, address *domain.Address) (error, *domain.Address)
	ListUserAddresses(userId string) (error, []*domain.Address)
}
