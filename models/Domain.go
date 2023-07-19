package models

import (
	"github.com/uptrace/bun"

	"github.com/google/uuid"
)

type Domain struct {
	bun.BaseModel `bun:"domains,alias:domains"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`

	Questions Questions `bun:"rel:has-many" json:"questions"`
}

type Domains []*Domain
