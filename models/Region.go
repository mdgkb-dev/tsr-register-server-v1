package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Region struct {
	bun.BaseModel `bun:"regions,alias:regions"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string    `json:"name"`
}

type Regions []*Region
