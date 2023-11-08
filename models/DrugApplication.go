package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugApplication struct {
	bun.BaseModel               `bun:"drug_applications,alias:drug_applications"`
	ID                          uuid.NullUUID               `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Date                        *time.Time                  `bun:"item_date" json:"date"`
	Number                      string                      `json:"number"`
	Status                      *Status                     `bun:"rel:belongs-to" json:"status"`
	StatusID                    uuid.NullUUID               `bun:"type:uuid" json:"statusId"`
	CommissionsDrugApplications CommissionsDrugApplications `bun:"rel:has-many" json:"commissionsDrugApplications"`
	DrugApplicationFiles        DrugApplicationFiles        `bun:"rel:has-many" json:"drugApplicationFiles"`
	FundContract                *FundContract               `bun:"rel:has-one" json:"fundContract"`
	//FundContractID              uuid.NullUUID               `bun:"type:uuid" json:"fundContractId"`
}

type DrugApplications []*DrugApplication
type DrugApplicationsWithCount struct {
	DrugApplications DrugApplications `json:"items"`
	Count            int              `json:"count"`
}
