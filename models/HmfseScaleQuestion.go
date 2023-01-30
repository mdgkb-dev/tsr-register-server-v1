package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type HmfseScaleQuestion struct {
	bun.BaseModel            `bun:"hmfse_scale_questions,alias:hmfse_scale_questions"`
	ID                       uuid.UUID                `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                     string                   `json:"name"`
	Order                    int                      `bun:"item_order" json:"order"`
	Description              string                   `json:"description"`
	HmfseScaleQuestionScores HmfseScaleQuestionScores `bun:"rel:has-many" json:"hmfseScaleQuestionScores"`
}

type HmfseScaleQuestions []*HmfseScaleQuestion
