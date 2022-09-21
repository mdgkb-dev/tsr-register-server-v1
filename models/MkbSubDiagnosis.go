package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MkbSubDiagnosis struct {
	bun.BaseModel        `bun:"mkb_sub_diagnosis_view,alias:mkb_sub_diagnosis_view"`
	ID                   uuid.NullUUID        `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                 string               `json:"name"`
	FullName             string               `json:"fullName"`
	SubCode              int                  `json:"subCode"`
	Comment              string               `json:"comment"`
	Leaf                 bool                 `json:"leaf"`
	Relevant             bool                 `json:"relevant"`
	MkbDiagnosis         *MkbDiagnosis        `bun:"rel:belongs-to" json:"mkbDiagnosis"`
	MkbDiagnosisID       uuid.UUID            `bun:"type:uuid" json:"mkbDiagnosisId"`
	MkbConcreteDiagnosis MkbConcreteDiagnoses `bun:"rel:has-many" json:"mkbConcreteDiagnosis"`
}

type MkbSubDiagnoses []*MkbSubDiagnosis
