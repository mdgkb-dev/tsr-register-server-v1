package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type CommissionDoctor struct {
	bun.BaseModel `bun:"commissions_doctors,alias:commissions_doctors"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	CommissionID  uuid.NullUUID `bun:"type:uuid" json:"commissionId"`
	Commission    *Commission   `bun:"rel:belongs-to" json:"commission"`
	DoctorID      uuid.NullUUID `bun:"type:uuid" json:"doctorId"`
	Doctor        *Doctor       `bun:"rel:belongs-to" json:"doctor"`
	Role          string        `json:"role"`
	Order         string        `bun:"item_order" json:"order"`
}

type CommissionsDoctors []*CommissionDoctor
type CommissionDoctorsWithCount struct {
	CommissionsDoctors CommissionsDoctors `json:"items"`
	Count              int                `json:"count"`
}
