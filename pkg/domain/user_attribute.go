package domain

import (
	"database/sql"
	db_commons "github.com/kutty-kumar/db_commons/model"
	"github.com/kutty-kumar/ho_oh/pikachu_v1"
)

type UserAttribute struct {
	db_commons.BaseDomain
	AttributeKey   string
	AttributeValue string
	UserID         string
}

func (i *UserAttribute) GetName() db_commons.DomainName {
	return "user_attributes"
}

func (i *UserAttribute) ToDto() interface{} {
	return pikachu_v1.UserAttributeDto{
		AttributeKey:   i.AttributeKey,
		AttributeValue: i.AttributeValue,
		ExternalId:     i.ExternalId,
	}
}

func (i *UserAttribute) FillProperties(dto interface{}) db_commons.Base {
	userAttributeDto := dto.(pikachu_v1.UserAttributeDto)
	i.AttributeKey = userAttributeDto.AttributeKey
	i.AttributeValue = userAttributeDto.AttributeValue
	return i
}

func (i *UserAttribute) Merge(other interface{}) {
	userAttributeDto := other.(*UserAttribute)
	if userAttributeDto.AttributeValue != "" {
		i.AttributeValue = userAttributeDto.AttributeValue
	}
}

func (i *UserAttribute) FromSqlRow(rows *sql.Rows) (db_commons.Base, error) {
	err := rows.Scan(&i.Id, &i.CreatedAt, &i.UpdatedAt, &i.DeletedAt, &i.Status, &i.AttributeKey, &i.AttributeValue)
	return i, err
}

func (i *UserAttribute) SetExternalId(externalId string) {
	i.ExternalId = externalId
}
