package domain

import (
	"database/sql"
	"github.com/kutty-kumar/db_commons/model"
	"pikachu/pkg/pb"
)

type Relation struct {
	db_commons.BaseDomain
	RelationType pb.Relation
	UserID       string
}

func (r *Relation) GetName() db_commons.DomainName {
	return "relations"
}

func (r *Relation) ToDto() interface{} {
	return pb.RelationDto{
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
