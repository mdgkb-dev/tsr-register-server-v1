package models

import (
	"github.com/uptrace/bun"

	"github.com/google/uuid"
)

type User struct {
	bun.BaseModel `bun:"users,alias:users"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Login         string        `json:"login"`

	RegisterPropertyToUser  RegisterPropertiesToUser `bun:"rel:has-many" json:"registerPropertyToUser"`
	RegistersUsers          RegistersUsers           `bun:"rel:has-many" json:"registersUsers"`
	RegistersUsersForDelete []uuid.UUID              `bun:"-" json:"registersUsersForDelete"`

	DomainID uuid.NullUUID `bun:"type:uuid" json:"-"`
	Domain   *Domain       `bun:"rel:belongs-to" json:"-"`

	UserAccountID uuid.NullUUID `bun:"type:uuid" json:"userAccountId"`
	UserAccount   *UserAccount  `bun:"rel:belongs-to" json:"userAccount"`
}

type Users []*User

type UsersWithCount struct {
	Users Users `json:"users"`
	Count int   `json:"count"`
}

func (item *User) SetIDForChildren() {
	for i := range item.RegistersUsers {
		item.RegistersUsers[i].UserID = item.ID.UUID
	}
}

type RegisterPropertyToUser struct {
	bun.BaseModel      `bun:"register_property_to_user,alias:register_property_to_user"`
	ID                 uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	RegisterProperty   *Domain   `bun:"rel:belongs-to" json:"registerProperty"`
	RegisterPropertyID uuid.UUID `bun:"type:uuid" json:"registerPropertyId"`
	User               *User     `bun:"rel:belongs-to" json:"user"`
	UserID             uuid.UUID `bun:"type:uuid" json:"userId"`
}

type RegisterPropertiesToUser []*RegisterPropertyToUser

func (item *User) SetJWTClaimsMap(claims map[string]interface{}) {
	claims["user_id"] = item.ID.UUID
	claims["user_domain_id"] = item.DomainID.UUID
}
