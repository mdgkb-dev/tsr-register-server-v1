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
	Register                                 *Register                          `bun:"rel:belongs-to" json:"register"`
	RegisterID                               uuid.UUID                          `bun:"type:uuid" json:"registerId"`
	RegisterQueryToRegisterProperty          []*RegisterQueryToRegisterProperty `bun:"rel:has-many" json:"registerQueryToRegisterProperty"`
	RegisterQueryToRegisterPropertyForDelete []string                           `bun:"-" json:"registerQueryToRegisterPropertyForDelete"`
}

type RegisterQueryToRegisterProperty struct {
	bun.BaseModel      `bun:"register_query_to_register_property,alias:register_query_to_register_property"`
	ID                 uuid.UUID         `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	RegisterQueryID    uuid.UUID         `bun:"type:uuid" json:"registerQueryId"`
	RegisterQuery      *RegisterQuery    `bun:"rel:belongs-to" json:"registerQuery"`
	RegisterPropertyID uuid.UUID         `bun:"type:uuid" json:"registerPropertyId"`
	RegisterProperty   *RegisterProperty `bun:"rel:belongs-to" json:"registerProperty"`
	Order              int               `bun:"order" json:"order"`
	IsAggregate        bool              `json:"isAggregate"`
}

func (query *RegisterQuery) SetIdForChildren() {
	if len(query.RegisterQueryToRegisterProperty) == 0 {
		return
	}

	for i := range query.RegisterQueryToRegisterProperty {
		query.RegisterQueryToRegisterProperty[i].RegisterQueryID = query.ID
	}
}
