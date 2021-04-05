package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kutty-kumar/db_commons/model"
)

type BaseDao struct {
	persistence      db_commons.BaseRepository
	db               *gorm.DB
	factory          db_commons.DomainFactory
	ExternalIdSetter func(externalId string, base db_commons.Base) db_commons.Base
}

func NewBaseDao(persistence db_commons.BaseRepository, db *gorm.DB, factory db_commons.DomainFactory,
	externalIdSetter func(externalId string, base db_commons.Base) db_commons.Base) BaseDao {
	return BaseDao{
		persistence:      persistence,
		db:               db,
		factory:          factory,
		ExternalIdSetter: externalIdSetter,
	}
}
