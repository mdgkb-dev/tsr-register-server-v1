package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type CustomSection struct {
	bun.BaseModel `bun:"custom_sections,alias:custom_sections"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string        `json:"name"`
	Component     string        `json:"component"`
	Order         int           `bun:"item_order" json:"order"`
}

type CustomSections []*CustomSection

type CustomSectionsWithCount struct {
	CustomSections CustomSections `json:"items"`
	Count          int            `json:"count"`
}
