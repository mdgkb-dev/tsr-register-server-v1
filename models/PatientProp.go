package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PatientProp struct {
	bun.BaseModel `bun:"patients,select:patients_view,alias:patients_view"`
	ModelInfo
	ID      uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Human   *Human        `bun:"rel:belongs-to" json:"human"`
	HumanID uuid.UUID     `bun:"type:uuid" json:"humanId"`
}

type PatientProps []*PatientProp

type PatientPropsWithCount struct {
	PatientProps PatientProps `json:"items"`
	Count        int          `json:"count"`
}
