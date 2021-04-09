package domain

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/kutty-kumar/db_commons/model"
	"github.com/kutty-kumar/ho_oh/core_v1"
	"github.com/kutty-kumar/ho_oh/pikachu_v1"
	"time"
)

type User struct {
	db_commons.BaseDomain
	FirstName      string
	LastName       string
	Gender         core_v1.Gender
	DateOfBirth    time.Time
	Identities     []Identity      `gorm:"foreignkey:UserID;association_foreignkey:ExternalId"`
	Addresses      []Address       `gorm:"foreignkey:UserID;association_foreignkey:ExternalId"`
	Relations      []Relation      `gorm:"foreignkey:UserID;association_foreignkey:ExternalId"`
	UserAttributes []UserAttribute `gorm:"foreignkey:UserID;association_foreignkey:ExternalId"`
	Age            int64
	Height         float64
	Weight         float64
}

func (u *User) GetName() db_commons.DomainName {
	return "users"
}

func (u *User) ToDto() interface{} {
	dobProto, _ := ptypes.TimestampProto(u.DateOfBirth)
	return &pikachu_v1.UserDto{
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Gender:      u.Gender,
		DateOfBirth: dobProto,
		Height:      u.Height,
		Weight:      u.Weight,
		Age:         u.Age,
		ExternalId:  u.ExternalId,
	}
}

func (u *User) FillProperties(dto interface{}) db_commons.Base {
	userDto := dto.(pikachu_v1.UserDto)
	u.FirstName = userDto.FirstName
	u.LastName = userDto.LastName
	u.Gender = userDto.Gender
	u.DateOfBirth = userDto.DateOfBirth.AsTime()
	u.Age = userDto.Age
	u.Height = userDto.Height
	u.Weight = userDto.Weight
	u.DeletedAt = nil
	return u
}

func (u *User) Merge(other interface{}) {
	updatableUser := other.(*User)
	if updatableUser.Age != 0 {
		u.Age = updatableUser.Age
	}
	if updatableUser.Weight != 0 {
		u.Weight = updatableUser.Weight
	}
	if updatableUser.Height != 0 {
		u.Height = updatableUser.Height
	}
	if updatableUser.FirstName != "" {
		u.FirstName = updatableUser.FirstName
	}
	if updatableUser.LastName != "" {
		u.LastName = updatableUser.LastName
	}
	if &updatableUser.DateOfBirth != nil {
		u.DateOfBirth = updatableUser.DateOfBirth
	}
}

func (u *User) FromSqlRow(rows *sql.Rows) (db_commons.Base, error) {
	err := rows.Scan(&u.ExternalId, &u.Id, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt, &u.Status, &u.FirstName,
		&u.LastName, &u.Gender, &u.DateOfBirth, &u.Age, &u.Height, &u.Weight)
	return u, err
}

func (u *User) AddIdentity(identity Identity) error {
	for _, existingIdentity := range u.Identities {
		if existingIdentity.IdentityType == identity.IdentityType && existingIdentity.IdentityValue == identity.IdentityValue {
			return errors.New(fmt.Sprintf("identity with value %v and type %v already exists", identity.IdentityType, identity.IdentityValue))
		}
	}
	u.Identities = append(u.Identities, identity)
	return nil
}

func (u *User) SetExternalId(externalId string) {
	u.ExternalId = externalId
}
