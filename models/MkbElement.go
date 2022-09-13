package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MkbElement struct {
	bun.BaseModel `bun:"mkb_flat_view,alias:mkb_flat_view"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	FullName      string    `json:"fullName"`
	Level         MkbLevel  `json:"level"`
	ClassID       uuid.UUID `json:"classId"`
}

type MkbLevel uint

const (
	MkbClassLevel MkbLevel = iota
	MkbGroupLevel
	MkbSubGroupLevel
	MkbSubSubGroupLevel
	MkbDiagnosisLevel
	MkbSubDiagnosisLevel
	MkbConcreteDiagnosisLevel
)
