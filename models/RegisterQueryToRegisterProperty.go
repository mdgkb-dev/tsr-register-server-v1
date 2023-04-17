package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterQueryToRegisterProperty struct {
	bun.BaseModel      `bun:"register_query_to_register_property,alias:register_query_to_register_property"`
	ID                 uuid.UUID      `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	RegisterQueryID    uuid.UUID      `bun:"type:uuid" json:"registerQueryId"`
	RegisterQuery      *RegisterQuery `bun:"rel:belongs-to" json:"registerQuery"`
	RegisterPropertyID uuid.UUID      `bun:"type:uuid" json:"registerPropertyId"`
	RegisterProperty   *Question      `bun:"rel:belongs-to" json:"registerProperty"`
	Order              int            `bun:"order" json:"order"`
	IsAggregate        bool           `json:"isAggregate"`
}

type RegisterQueryToRegisterProperties []*RegisterQueryToRegisterProperty
