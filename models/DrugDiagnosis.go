package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugDiagnosis struct {
	bun.BaseModel     `bun:"drugs_diagnosis,alias:drugs_diagnosis"`
	ID                uuid.UUID        `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Drug              *Drug            `bun:"rel:has-one" json:"drugs"`
	DrugID            uuid.UUID        `bun:"type:uuid" json:"drugId"`
	MkbDiagnosis      *MkbDiagnosis    `bun:"rel:belongs-to" json:"mkbDiagnosis"`
	MkbDiagnosisID    uuid.UUID        `bun:"type:uuid" json:"mkbDiagnosisId"`
	MkbSubDiagnosis   *MkbSubDiagnosis `bun:"rel:belongs-to" json:"mkbSubDiagnosis"`
	MkbSubDiagnosisID uuid.NullUUID    `bun:"type:uuid,nullzero" json:"mkbSubDiagnosisId"`
}

type DrugsDiagnosis []*DrugDiagnosis
