package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type CommissionTemplate struct {
	bun.BaseModel `bun:"commissions_templates,alias:commissions_templates"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Volume        string        `json:"volume"`
	Name          string        `json:"name"`
	DrugRegimen   *DrugRegimen  `bun:"rel:belongs-to" json:"drugRegimen"`
	DrugRegimenID uuid.NullUUID `bun:"type:uuid" json:"drugRegimenId"`

	Drug   *Drug         `bun:"rel:belongs-to" json:"drug"`
	DrugID uuid.NullUUID `bun:"type:uuid" json:"drugId"`

	CommissionsDoctorsTemplates CommissionsDoctorsTemplates `bun:"rel:has-many" json:"commissionsDoctorsTemplates"`
}

type CommissionsTemplates []*CommissionTemplate
type CommissionTemplatesWithCount struct {
	CommissionsTemplates CommissionsTemplates `json:"items"`
	Count                int                  `json:"count"`
}
