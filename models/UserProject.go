package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserProject struct {
	bun.BaseModel `bun:"users_projects,alias:users_projects"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	DomainID uuid.NullUUID `bun:"type:uuid" json:"-"`
	Domain   *Domain       `bun:"rel:belongs-to" json:"-"`

	UserID uuid.NullUUID `bun:"type:uuid" json:"-"`
	User   *User         `bun:"rel:belongs-to" json:"-"`
}
type UsersProjects []*UserProject
