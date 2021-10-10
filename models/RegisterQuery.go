package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterQuery struct {
	bun.BaseModel                            `bun:"register_queries,alias:register_queries"`
	ID                                       uuid.UUID                          `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Name                                     string                             `json:"name"`
	Type                                     string                             `bun:"type:register_query_type_enum" json:"type"`
	RegisterID                               uuid.UUID                          `bun:"type:uuid" json:"registerId"`
	RegisterQueryToRegisterProperty          []*RegisterQueryToRegisterProperty `bun:"rel:has-many" json:"registerQueryToRegisterProperty"`
	RegisterQueryToRegisterPropertyForDelete []string                           `bun:"-" json:"registerQueryToRegisterPropertyForDelete"`
}

type RegisterQueries []*RegisterQuery

func (query *RegisterQuery) SetIdForChildren() {
	if len(query.RegisterQueryToRegisterProperty) == 0 {
		return
	}

	for i := range query.RegisterQueryToRegisterProperty {
		query.RegisterQueryToRegisterProperty[i].RegisterQueryID = query.ID
	}
}
