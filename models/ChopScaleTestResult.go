package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ChopScaleTestResult struct {
	bun.BaseModel            `bun:"chop_scale_test_results,alias:chop_scale_test_results"`
	ID                       uuid.UUID               `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	ChopScaleTest            *ChopScaleTest          `bun:"rel:belongs-to" json:"chopScaleTest"`
	ChopScaleTestID          uuid.UUID               `bun:"type:uuid" json:"chopScaleTestId"`
	ChopScaleQuestionScore   *ChopScaleQuestionScore `bun:"rel:belongs-to" json:"chopScaleQuestionScore"`
	ChopScaleQuestionScoreID uuid.UUID               `bun:"type:uuid" json:"chopScaleQuestionScoreId"`
}
type ChopScaleTestResults []*ChopScaleTestResult
