package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type CommissionDoctorTemplate struct {
	bun.BaseModel        `bun:"commissions_doctors_templates,alias:commissions_doctors_templates"`
	ID                   uuid.NullUUID       `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	CommissionTemplate   *CommissionTemplate `bun:"rel:belongs-to" json:"commissionTemplate"`
	CommissionTemplateID uuid.NullUUID       `bun:"type:uuid" json:"commissionTemplateId"`
	DoctorID             uuid.NullUUID       `bun:"type:uuid" json:"doctorId"`
	Doctor               *Doctor             `bun:"rel:belongs-to" json:"doctor"`
	Role                 string              `json:"role"`
	Order                string              `bun:"item_order" json:"order"`
}

type CommissionsDoctorsTemplates []*CommissionDoctorTemplate
type CommissionsDoctorsTemplatesWithCount struct {
	CommissionsDoctorsTemplates CommissionsDoctorsTemplates `json:"items"`
	Count                       int                         `json:"count"`
}
