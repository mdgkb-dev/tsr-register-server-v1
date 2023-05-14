package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugApplication struct {
	bun.BaseModel `bun:"drug_applications,alias:drug_applications"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Date          *time.Time    `bun:"item_date" json:"name"`
	Commission    *Commission   `bun:"rel:belongs-to" json:"commission"`
	CommissionID  uuid.NullUUID `bun:"type:uuid" json:"commissionId"`
}

type DrugApplications []*DrugApplication
type DrugApplicationsWithCount struct {
	DrugApplications DrugApplications `json:"items"`
	Count            int              `json:"count"`
}
