package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MkbConcreteDiagnosis struct {
	bun.BaseModel     `bun:"mkb_concrete_diagnosis,alias:mkb_concrete_diagnosis"`
	ID                uuid.UUID        `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name              string           `json:"name"`
	Comment           string           `json:"comment"`
	Leaf              bool             `json:"bool"`
	Relevant          bool             `json:"relevant"`
	MkbSubDiagnosis   *MkbSubDiagnosis `bun:"rel:belongs-to" json:"mkbSubDiagnosis"`
	MkbSubDiagnosisID uuid.UUID        `bun:"type:uuid" json:"mkbSubDiagnosisId"`
}

type MkbConcreteDiagnoses []*MkbConcreteDiagnosis
