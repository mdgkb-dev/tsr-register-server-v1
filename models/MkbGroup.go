package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MkbGroup struct {
	bun.BaseModel `bun:"mkb_group,alias:mkb_group"`
	ID            uuid.UUID       `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string          `json:"name"`
	RangeStart    string          `json:"rangeStart"`
	RangeEnd      string          `json:"rangeEnd"`
	Comment       string          `json:"comment"`
	Leaf          bool            `json:"bool"`
	Relevant      bool            `json:"relevant"`
	MkbClass      *MkbClass       `bun:"rel:belongs-to" json:"mkbClass"`
	MkbClassID    uuid.UUID       `bun:"type:uuid" json:"mkbClassId"`
	MkbSubGroups  []*MkbSubGroup  `bun:"rel:has-many" json:"mkbSubGroups"`
	MkbDiagnosis  []*MkbDiagnosis `bun:"rel:has-many"json:"mkbDiagnosis"`
}
