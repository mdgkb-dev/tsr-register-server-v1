package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type HmfseScaleTestResult struct {
	bun.BaseModel             `bun:"hmfse_scale_test_results,alias:hmfse_scale_test_results"`
	ID                        uuid.UUID                `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Date                      time.Time                `bun:"item_date" json:"date"`
	HmfseScaleTest            *HmfseScaleTest          `bun:"rel:belongs-to" json:"hmfseScaleTest"`
	HmfseScaleTestID          uuid.UUID                `bun:"type:uuid" json:"hmfseScaleTestId"`
	HmfseScaleQuestionScore   *HmfseScaleQuestionScore `bun:"rel:belongs-to" json:"hmfseScaleQuestionScore"`
	HmfseScaleQuestionScoreID uuid.UUID                `bun:"type:uuid" json:"hmfseScaleQuestionScoreId"`
}
type HmfseScaleTestResults []*HmfseScaleTestResult
