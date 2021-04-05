package repository

import (
	"errors"
	"pikachu/pkg/domain"
	"pikachu/pkg/pb"
)

type IdentityGORMRepository struct {
	BaseDao
}

func (ir *IdentityGORMRepository) GetIdentity(userId string, identityId string, iType pb.IdentityType) (error, *domain.User, *domain.Identity) {
	err, user := ir.persistence.GetByExternalId(userId, ir.factory.GetMapping("user"))
	if err != nil {
		return err, nil, nil
	}
	nIdentity := interface{}(ir.factory.GetMapping("identity")()).(*domain.Identity)
	if err := ir.db.Model(nIdentity).Where("user_id = ? AND external_id = ? AND identity_type = ?", userId, identityId, iType).Find(nIdentity).Error; err != nil {
		return err, user.(*domain.User), nil
	}
	return nil, interface{}(user).(*domain.User), interface{}(nIdentity).(*domain.Identity)

}

func (ir *IdentityGORMRepository) GetExistingIdentity(userId string, identity *domain.Identity) (error, *domain.User, *domain.Identity) {
	err, user := ir.persistence.GetByExternalId(userId, ir.factory.GetMapping("user"))
	if err != nil {
		return err, nil, nil
	}
	nIdentity := interface{}(ir.factory.GetMapping("identity")()).(*domain.Identity)
	if err := ir.db.Model(identity).Where("user_id = ? AND identity_type = ? AND identity_value = ?", user.GetId(), identity.IdentityType, identity.IdentityValue).Find(nIdentity).Error; err != nil {
		return err, user.(*domain.User), nil
	}
	return nil, interface{}(user).(*domain.User), interface{}(nIdentity).(*domain.Identity)

}

func (ir *IdentityGORMRepository) CreateIdentity(userId string, identity *domain.Identity) (error, *domain.Identity) {
	err, user, eIdentity := ir.GetExistingIdentity(userId, identity)
	if (err != nil && user == nil && eIdentity == nil) || (err == nil && user != nil && eIdentity != nil) {
		return errors.New("either identity exists or user doesn't exist"), nil
	}

	identity.UserID = user.GetExternalId()
	err, cIdentity := ir.persistence.Create(identity, ir.ExternalIdSetter)
	if err != nil {
		return err, nil
	}
	return nil, cIdentity.(*domain.Identity)
}

func (ir *IdentityGORMRepository) UpdateIdentity(userId string, identityId string, identity *domain.Identity) (error, *domain.Identity) {
	err, _, eIdentity := ir.GetIdentity(userId, identityId, identity.IdentityType)
	if err != nil || eIdentity == nil {
		return errors.New("either identity doesn't exists or user doesn't exist"), nil
	}
	eIdentity.Merge(identity)
	err, uIdentity := ir.persistence.Update(identityId, eIdentity, ir.factory.GetMapping("identity"))
	if err != nil {
		return err, nil
	}
	return nil, interface{}(uIdentity).(*domain.Identity)
}

func (ir *IdentityGORMRepository) ListIdentities(userId string) (error, []domain.Identity) {
	user := ir.factory.GetMapping("user")().(*domain.User)
	user.ExternalId = userId
	identities := make([]domain.Identity, 0)
	if err := ir.db.Model(user).Where("external_id = ?", userId).Find(user).Error; err != nil {
		return err, nil
	}
	if err := ir.db.Model(user).Related(&identities).Error; err != nil {
		return err, nil
	}
	return nil, identities
}

func NewIdentityGormRepository(dao BaseDao) IdentityGORMRepository {
	return IdentityGORMRepository{
		dao,
	}
}
