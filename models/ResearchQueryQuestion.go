package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResearchQueryQuestion struct {
	bun.BaseModel   `bun:"register_query_to_register_property,alias:register_query_to_register_property"`
	ID              uuid.UUID      `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	ResearchQueryID uuid.UUID      `bun:"type:uuid" json:"researchQueryId"`
	ResearchQuery   *ResearchQuery `bun:"rel:belongs-to" json:"researchQuery"`
	QuestionID      uuid.UUID      `bun:"type:uuid" json:"questionId"`
	Question        *Question      `bun:"rel:belongs-to" json:"question"`
	Order           int            `bun:"order" json:"order"`
	IsAggregate     bool           `json:"isAggregate"`
}

type ResearchQueriesQuestions []*ResearchQueryQuestion
