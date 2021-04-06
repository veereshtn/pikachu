package repository

import (
	"errors"
	"github.com/kutty-kumar/db_commons/model"
	"pikachu/pkg/domain"
)

type UserGormRepository struct {
	db_commons.BaseDao
}

func NewUserGormRepository(dao db_commons.BaseDao) UserGormRepository {
	return UserGormRepository{
		dao,
	}
}

func (u *UserGormRepository) Create(user *domain.User) (error, *domain.User) {
	err, base := u.Create(user)
	return err, interface{}(base).(*domain.User)
}

func (u *UserGormRepository) Update(id string, user *domain.User) (error, *domain.User) {
	err, base := u.Update(id, user)
	if err != nil {
		return err, nil
	}
	return err, interface{}(base).(*domain.User)
}

func (u *UserGormRepository) FindByExternalId(id string) (error, *domain.User) {
	err, base := u.GetByExternalId(id)
	if base != nil {
		return err, base.(*domain.User)
	}
	return errors.New("not found"), nil
}

func (u *UserGormRepository) MultiGetByExternalIds(ids []string) (error, []domain.User) {
	var userSlice []domain.User
	err, sqlSlice := u.MultiGetByExternalId(ids)
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
