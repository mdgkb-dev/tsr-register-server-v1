package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ChopScaleQuestion struct {
	bun.BaseModel           `bun:"chop_scale_questions,alias:chop_scale_questions"`
	ID                      uuid.UUID               `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                    string                  `json:"name"`
	Order                   int                     `bun:"item_order" json:"order"`
	Description             string                  `json:"description"`
	ChopScaleQuestionScores ChopScaleQuestionScores `bun:"rel:has-many" json:"chopScaleQuestionScores"`
}

type ChopScaleQuestions []*ChopScaleQuestion
