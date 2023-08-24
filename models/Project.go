package models

import (
	"github.com/uptrace/bun"

	"github.com/google/uuid"
)

type Project struct {
	bun.BaseModel `bun:"users,alias:users"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`

	DomainID uuid.NullUUID `bun:"type:uuid" json:"-"`
	Domain   *Domain       `bun:"rel:belongs-to" json:"-"`
}

type Projects []*Project

type ProjectsWithCount struct {
	Projects Projects `json:"users"`
	Count    int      `json:"count"`
}
