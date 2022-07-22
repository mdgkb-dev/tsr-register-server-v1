package models

import (
	"github.com/uptrace/bun"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel          `bun:"users,alias:users"`
	ID                     uuid.UUID                `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Login                  string                   `json:"login"`
	Password               string                   `json:"password"`
	RegisterPropertyToUser RegisterPropertiesToUser `bun:"rel:has-many" json:"registerPropertyToUser"`
	//
	RegistersUsers          RegistersUsers `bun:"rel:has-many" json:"registersUsers"`
	RegistersUsersForDelete []uuid.UUID    `bun:"-" json:"registersUsersForDelete"`
}

type Users []*User

func (item *User) SetIdForChildren() {
	for i := range item.RegistersUsers {
		item.RegistersUsers[i].UserID = item.ID
	}
}

type RegisterPropertyToUser struct {
	bun.BaseModel      `bun:"register_property_to_user,alias:register_property_to_user"`
	ID                 uuid.UUID         `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	RegisterProperty   *RegisterProperty `bun:"rel:belongs-to" json:"registerProperty"`
	RegisterPropertyID uuid.UUID         `bun:"type:uuid" json:"registerPropertyId"`
	User               *User             `bun:"rel:belongs-to" json:"user"`
	UserID             uuid.UUID         `bun:"type:uuid" json:"userId"`
}

type RegisterPropertiesToUser []*RegisterPropertyToUser

func (i *User) GenerateHashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(i.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	i.Password = string(hash)
	return nil
}

func (i *User) CompareWithHashPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(i.Password), []byte(password)) == nil
}
