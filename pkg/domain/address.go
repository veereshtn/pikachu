package domain

import (
	"database/sql"
	"errors"
	"github.com/kutty-kumar/db_commons/model"
	"github.com/kutty-kumar/ho_oh/pkg/pikachu_v1"
)

type Address struct {
	db_commons.BaseDomain
	Line1   string
	Line2   string
	ZipCode string
	State   string
	Country string
	UserID  string
}

func (a *Address) GetName() db_commons.DomainName {
	return "addresses"
}

func (a *Address) ToDto() interface{} {
	return pikachu_v1.AddressDto{
		Country: a.Country,
		Line1:   a.Line1,
		Line2:   a.Line2,
		State:   a.State,
		ZipCode: a.ZipCode,
	}
}

func (a *Address) FillProperties(dto interface{}) db_commons.Base {
	addressDto := dto.(pikachu_v1.AddressDto)
	a.Country = addressDto.Country
	a.Line2 = addressDto.Line2
	a.Line1 = addressDto.Line1
	a.State = addressDto.State
	a.ZipCode = addressDto.State
	return a
}

func (a *Address) Merge(other interface{}) {
	address := other.(pikachu_v1.AddressDto)
	if address.Line1 != "" {
		a.Line1 = address.Line1
	}
	if address.Line2 != "" {
		a.Line2 = address.Line2
	}
	if address.Country != "" {
		a.Country = address.Country
	}
	if address.State != "" {
		a.State = address.State
	}
	if address.ZipCode != "" {
		a.ZipCode = address.ZipCode
	}
}

func (a *Address) FromSqlRow(rows *sql.Rows) (db_commons.Base, error) {
	return nil, errors.New("not implemented")
}

func (a *Address) SetExternalId(externalId string) {
	a.ExternalId = externalId
}
