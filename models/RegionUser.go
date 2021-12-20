package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegionUser struct {
	bun.BaseModel      `bun:"regions_users,alias:regions_users"`
	ID                 uuid.UUID         `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Region   *Region `bun:"rel:belongs-to" json:"region"`
	RegionID uuid.UUID         `bun:"type:uuid" json:"regionId"`
	User               *User             `bun:"rel:belongs-to" json:"user"`
	UserID             uuid.UUID         `bun:"type:uuid" json:"userId"`
}

type RegionsUsers []*RegionUser
