package repository

import (
	"errors"
	"github.com/kutty-kumar/db_commons/model"
	"pikachu/pkg/domain"
)

type IdentityGORMRepository struct {
	db_commons.BaseDao
}

func (ir *IdentityGORMRepository) GetIdentity(userId string, identity *domain.Identity) (error, *domain.Identity) {
	nIdentity := &domain.Identity{}
	if err := ir.GetDb().Model(nIdentity).Where("user_id = ? AND external_id = ? AND identity_type = ?", userId, identity.GetExternalId(), identity.IdentityType).Find(nIdentity).Error; err != nil {
		return err, nil
	}
	return nil, interface{}(nIdentity).(*domain.Identity)

}

func (ir *IdentityGORMRepository) GetExistingIdentity(userId string, identity *domain.Identity) (error, *domain.Identity) {
	err, user := ir.GetByExternalId(userId)
	if err != nil {
		return err, nil
	}
	nIdentity := &domain.Identity{}
	if err := ir.GetDb().Model(identity).Where("user_id = ? AND identity_type = ? AND identity_value = ?", user.GetId(), identity.IdentityType, identity.IdentityValue).Find(nIdentity).Error; err != nil {
		return err, nil
	}
	return nil, interface{}(nIdentity).(*domain.Identity)

}

func (ir *IdentityGORMRepository) CreateIdentity(userId string, identity *domain.Identity) (error, *domain.Identity) {
	err, eIdentity := ir.GetExistingIdentity(userId, identity)
	if (err != nil && eIdentity == nil) || (err == nil && eIdentity != nil) {
		return errors.New("either identity exists or user doesn't exist"), nil
	}
	identity.UserID = userId
	err, cIdentity := ir.Create(identity)
	if err != nil {
		return err, nil
	}
	return nil, cIdentity.(*domain.Identity)
}

func (ir *IdentityGORMRepository) UpdateIdentity(userId string, identityId string, identity *domain.Identity) (error, *domain.Identity) {
	err, eIdentity := ir.GetIdentity(userId, identity)
	if err != nil || eIdentity == nil {
		return errors.New("either identity doesn't exists or user doesn't exist"), nil
	}
	eIdentity.Merge(identity)
	err, uIdentity := ir.Update(identityId, eIdentity)
	if err != nil {
		return err, nil
	}
	return nil, interface{}(uIdentity).(*domain.Identity)
}

func (ir *IdentityGORMRepository) ListIdentities(userId string) (error, []domain.Identity) {
	user := domain.User{}
	user.ExternalId = userId
	identities := make([]domain.Identity, 0)
	if err := ir.GetDb().Model(user).Where("external_id = ?", userId).Find(user).Error; err != nil {
		return err, nil
	}
	if err := ir.GetDb().Model(user).Related(&identities).Error; err != nil {
		return err, nil
	}
	return nil, identities
}

func NewIdentityGormRepository(dao db_commons.BaseDao) IdentityGORMRepository {
	return IdentityGORMRepository{
		dao,
	}
}
