package repository

import (
	"pikachu/pkg/domain"
)

type AddressGORMRepository struct {
	BaseDao
}

func (ar *AddressGORMRepository) CreateUserAddress(userId string, address *domain.Address) (error, *domain.Address) {
	err, user := ar.persistence.GetByExternalId(userId, ar.factory.GetMapping("user"))
	if err != nil {
		return err, nil
	}
	address.UserID = user.GetExternalId()
	err, bAddress := ar.persistence.Create(address, ar.ExternalIdSetter)
	if err != nil {
		return err, nil
	}
	return nil, bAddress.(*domain.Address)
}

func (ar *AddressGORMRepository) UpdateUserAddress(userId string, addressId string, address *domain.Address) (error, *domain.Address) {
	err, user := ar.persistence.GetByExternalId(userId, ar.factory.GetMapping("user"))
	if err != nil {
		return err, nil
	}
	eAddress := interface{}(ar.factory.GetMapping("address")()).(domain.Address)
	if err := ar.db.Model(&eAddress).Where("external_id = ? AND user_id = ?", addressId, user.GetId()).Find(&eAddress).Error; err != nil {
		return err, nil
	}
	eAddress.Merge(address)
	err, uAddress := ar.persistence.Update(addressId, &eAddress, ar.factory.GetMapping("address"))
	if err != nil {
		return err, nil
	}
	return nil, interface{}(uAddress).(*domain.Address)
}

func (ar *AddressGORMRepository) ListUserAddresses(userId string) (error, []*domain.Address) {
	addresses := make([]*domain.Address, 0)
	err, user := ar.persistence.GetByExternalId(userId, ar.factory.GetMapping("user"))
	if err != nil {
		return err, nil
	}
	if err := ar.db.Model(user).Related(&addresses).Error; err != nil {
		return err, nil
	}
	return nil, addresses
}
