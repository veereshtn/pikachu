package repository

import (
	"errors"
	db_commons "github.com/kutty-kumar/db_commons/model"
	"pikachu/pkg/domain"
)

type UserAttributeGORMRepository struct {
	db_commons.BaseDao
}

func (ir *UserAttributeGORMRepository) GetUserAttribute(userId string, userAttribute *domain.UserAttribute) (error, *domain.UserAttribute) {
	nUserAttribute := &domain.UserAttribute{}
	if err := ir.GetDb().Model(nUserAttribute).Where("user_id = ? AND attribute_key = ?", userId, userAttribute.AttributeKey).Find(nUserAttribute).Error; err != nil {
		return err, nil
	}
	return nil, interface{}(nUserAttribute).(*domain.UserAttribute)
}

func (ir *UserAttributeGORMRepository) CreateUserAttribute(userId string, userAttribute *domain.UserAttribute) (error, *domain.UserAttribute) {
	err, eUserAttribute := ir.GetUserAttribute(userId, userAttribute)
	if (err != nil && eUserAttribute == nil) || (err == nil && eUserAttribute != nil) {
		return errors.New("either user attribute exists or user doesn't exist"), nil
	}
	userAttribute.UserID = userId
	err, cUserAttribute := ir.Create(userAttribute)
	if err != nil {
		return err, nil
	}
	return nil, cUserAttribute.(*domain.UserAttribute)
}

func (ir *UserAttributeGORMRepository) UpdateUserAttribute(userId string, userAttribute *domain.UserAttribute) (error, *domain.UserAttribute) {
	err, eUserAttribute := ir.GetUserAttribute(userId, userAttribute)
	if err != nil || eUserAttribute == nil {
		return errors.New("either user attribute doesn't exists or user doesn't exist"), nil
	}
	eUserAttribute.Merge(userAttribute)
	err, uUserAttribute := ir.Update(eUserAttribute.ExternalId, eUserAttribute)
	if err != nil {
		return err, nil
	}
	return nil, interface{}(uUserAttribute).(*domain.UserAttribute)
}

func (ir *UserAttributeGORMRepository) ListUserAttributes(userId string) (error, []domain.UserAttribute) {
	user := domain.User{}
	user.ExternalId = userId
	userAttributes := make([]domain.UserAttribute, 0)
	if err := ir.GetDb().Model(user).Where("external_id = ?", userId).Find(user).Error; err != nil {
		return err, nil
	}
	if err := ir.GetDb().Model(user).Related(&userAttributes).Error; err != nil {
		return err, nil
	}
	return nil, userAttributes
}

func NewUserAttributeGormRepository(dao db_commons.BaseDao) UserAttributeGORMRepository {
	return UserAttributeGORMRepository{
		dao,
	}
}
