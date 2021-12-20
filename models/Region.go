package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Region struct {
	bun.BaseModel                            `bun:"region,alias:region"`
	ID                                       uuid.UUID                          `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Name                                     string                             `json:"name"`
}

type Regions []*Region
