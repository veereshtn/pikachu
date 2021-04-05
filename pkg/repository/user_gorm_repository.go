package repository

import (
	"errors"
	"pikachu/pkg/domain"
)

type UserGormRepository struct {
	BaseDao
}

func NewUserGormRepository(dao BaseDao) UserGormRepository {
	return UserGormRepository{
		dao,
	}
}

func (u *UserGormRepository) Create(user *domain.User) (error, *domain.User) {
	err, base := u.persistence.Create(user, u.ExternalIdSetter)
	return err, base.(*domain.User)
}

func (u *UserGormRepository) Update(id string, user *domain.User) (error, *domain.User) {
	err, base := u.persistence.Update(id, user, u.factory.GetMapping("user"))
	if err != nil {
		return err, nil
	}
	return err, base.(*domain.User)
}

func (u *UserGormRepository) FindByExternalId(id string) (error, *domain.User) {
	err, base := u.persistence.GetByExternalId(id, u.factory.GetMapping("user"))
	if base != nil {
		return err, base.(*domain.User)
	}
	return errors.New("not found"), nil
}

func (u *UserGormRepository) MultiGetByExternalIds(ids []string) (error, []domain.User) {
	var userSlice []domain.User
	err, sqlSlice := u.persistence.MultiGetByExternalId(ids, u.factory.GetMapping("user"))
	if err != nil {
		return err, nil
	}

	for _, sqlRow := range sqlSlice {
		userSlice = append(userSlice, *interface{}(sqlRow).(*domain.User))
	}
	return nil, userSlice
}

func (u *UserGormRepository) handleError(user *domain.User, err error) (error, *domain.User) {
	if err != nil {
		return err, nil
	}
	return nil, user
}

func (u *UserGormRepository) UpdateIdentity(userId string, identityId string, identity *domain.Identity) (error, *domain.Identity) {
	panic("implement me")
}

func (u *UserGormRepository) ListIdentities(userId string) (error, []domain.Identity) {
	panic("implement me")
}
