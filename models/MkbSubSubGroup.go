package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MkbSubSubGroup struct {
	bun.BaseModel `bun:"mkb_class,alias:mkb_class"`
	ID            uuid.UUID       `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string          `json:"name"`
	RangeStart    string          `json:"rangeStart"`
	RangeEnd      string          `json:"rangeEnd"`
	Comment       string          `json:"comment"`
	Leaf          bool            `json:"bool"`
	Relevant      bool            `json:"relevant"`
	MkbSubGroup   *MkbSubGroup    `bun:"rel:belongs-to" json:"mkbSubGroup"`
	MkbSubGroupID uuid.UUID       `bun:"type:uuid" json:"mkbSubGroupId"`
	MkbDiagnosis  []*MkbDiagnosis `bun:"rel:has-many"json:"mkbDiagnosis"`
}
