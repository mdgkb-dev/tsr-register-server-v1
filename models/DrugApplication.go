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
	//Commissions   Commissions   `bun:"rel:belongs-to" json:"commission"`
	Number                      string                      `json:"number"`
	DrugApplicationStatus       *DrugApplicationStatus      `bun:"rel:belongs-to" json:"drugApplicationStatus"`
	DrugApplicationStatusID     uuid.NullUUID               `bun:"type:uuid" json:"drugApplicationStatusId"`
	CommissionsDrugApplications CommissionsDrugApplications `bun:"rel:has-many" json:"commissionsDrugApplications"`
}

type DrugApplications []*DrugApplication
type DrugApplicationsWithCount struct {
	DrugApplications DrugApplications `json:"items"`
	Count            int              `json:"count"`
}
