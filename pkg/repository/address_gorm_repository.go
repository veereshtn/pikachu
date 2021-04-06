package repository

import (
	"github.com/kutty-kumar/db_commons/model"
	"pikachu/pkg/domain"
)

type AddressGORMRepository struct {
	db_commons.BaseRepository
}

func (ar *AddressGORMRepository) CreateUserAddress(userId string, address *domain.Address) (error, *domain.Address) {
	err, user := ar.GetByExternalId(userId)
	if err != nil {
		return err, nil
	}
	address.UserID = user.GetExternalId()
	err, bAddress := ar.Create(address)
	if err != nil {
		return err, nil
	}
	return nil, bAddress.(*domain.Address)
}

func (ar *AddressGORMRepository) UpdateUserAddress(userId string, addressId string, address *domain.Address) (error, *domain.Address) {
	err, user := ar.GetByExternalId(userId)
	if err != nil {
		return err, nil
	}
	eAddress := domain.Address{}
	if err := ar.GetDb().Model(&eAddress).Where("external_id = ? AND user_id = ?", addressId, user.GetId()).Find(&eAddress).Error; err != nil {
		return err, nil
	}
	eAddress.Merge(address)
	err, uAddress := ar.Update(addressId, &eAddress)
	if err != nil {
		return err, nil
	}
	return nil, interface{}(uAddress).(*domain.Address)
}

func (ar *AddressGORMRepository) ListUserAddresses(userId string) (error, []*domain.Address) {
	addresses := make([]*domain.Address, 0)
	err, user := ar.GetByExternalId(userId)
	if err != nil {
		return err, nil
	}
	if err := ar.GetDb().Model(user).Related(&addresses).Error; err != nil {
		return err, nil
	}
	return nil, addresses
}
