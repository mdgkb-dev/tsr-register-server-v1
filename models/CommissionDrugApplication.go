package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type CommissionDrugApplication struct {
	bun.BaseModel     `bun:"commissions_drug_applications,alias:commissions_drug_applications"`
	ID                uuid.NullUUID    `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Commission        *Commission      `bun:"rel:belongs-to" json:"commission"`
	CommissionID      uuid.NullUUID    `bun:"type:uuid" json:"commissionId"`
	DrugApplication   *DrugApplication `bun:"rel:belongs-to" json:"drugApplication"`
	DrugApplicationID uuid.NullUUID    `bun:"type:uuid" json:"drugApplicationId"`
}

type CommissionsDrugApplications []*CommissionDrugApplication
type CommissionsDrugApplicationsWithCount struct {
	CommissionsDrugApplications CommissionsDrugApplications `json:"items"`
	Count                       int                         `json:"count"`
}
