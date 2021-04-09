package domain

import (
	"database/sql"
	"github.com/kutty-kumar/db_commons/model"
	"github.com/kutty-kumar/ho_oh/core_v1"
	"github.com/kutty-kumar/ho_oh/pikachu_v1"
)

type Identity struct {
	db_commons.BaseDomain
	IdentityType  core_v1.IdentityType
	IdentityValue string
	UserID        string
}

func (i *Identity) GetName() db_commons.DomainName {
	return "identities"
}

func (i *Identity) ToDto() interface{} {
	return pikachu_v1.IdentityDto{
		IdentityType:  i.IdentityType,
		IdentityValue: i.IdentityValue,
		ExternalId:    i.ExternalId,
	}
}

func (i *Identity) FillProperties(dto interface{}) db_commons.Base {
	identityDto := dto.(pikachu_v1.IdentityDto)
	i.IdentityType = identityDto.IdentityType
	i.IdentityValue = identityDto.IdentityValue
	return i
}

func (i *Identity) Merge(other interface{}) {
	identityDto := other.(*Identity)
	if identityDto.IdentityType != 0 {
		i.IdentityType = identityDto.IdentityType
	}
	if identityDto.IdentityValue != "" {
		i.IdentityValue = identityDto.IdentityValue
	}
}

func (i *Identity) FromSqlRow(rows *sql.Rows) (db_commons.Base, error) {
	err := rows.Scan(&i.Id, &i.CreatedAt, &i.UpdatedAt, &i.DeletedAt, &i.Status, &i.IdentityType, &i.IdentityValue)
	return i, err
}

func (i *Identity) SetExternalId(externalId string) {
	i.ExternalId = externalId
}
