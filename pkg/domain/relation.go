package domain

import (
	"database/sql"
	"github.com/kutty-kumar/db_commons/model"
	"github.com/kutty-kumar/ho_oh/pkg/core_v1"
	"github.com/kutty-kumar/ho_oh/pkg/pikachu_v1"
)

type Relation struct {
	db_commons.BaseDomain
	RelationType core_v1.Relation
	UserID       string
}

func (r *Relation) GetName() db_commons.DomainName {
	return "relations"
}

func (r *Relation) ToDto() interface{} {
	return pikachu_v1.RelationDto{
	}
}

func (r *Relation) FillProperties(dto interface{}) db_commons.Base {
	panic("implement me")
}

func (r *Relation) Merge(other interface{}) {
	panic("implement me")
}

func (r *Relation) FromSqlRow(rows *sql.Rows) (db_commons.Base, error) {
	panic("implement me")
}

func (r *Relation) SetExternalId(externalId string) {
	panic("implement me")
}
