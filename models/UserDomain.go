package models

import (
	"github.com/uptrace/bun"

	"github.com/google/uuid"
)

type UserDomain struct {
	bun.BaseModel `bun:"users_domains,alias:users_domains"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	DomainID uuid.NullUUID `bun:"type:uuid" json:"-"`
	Domain   *Domain       `bun:"rel:belongs-to" json:"-"`

	UserID uuid.NullUUID `bun:"type:uuid" json:"userId"`
	User   *Domain       `bun:"rel:belongs-to" json:"user"`
}

type UsersDomains []*UserDomain

type UsersDomainsWithCount struct {
	UsersDomains UsersDomains `json:"users"`
	Count        int          `json:"count"`
}
