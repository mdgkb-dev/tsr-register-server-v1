package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MkbGroup struct {
	bun.BaseModel `bun:"mkb_groups_view,alias:mkb_groups_view"`
	ID            uuid.UUID    `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string       `json:"name"`
	FullName      string       `json:"fullName"`
	RangeStart    string       `json:"rangeStart"`
	RangeEnd      string       `json:"rangeEnd"`
	Comment       string       `json:"comment"`
	Leaf          bool         `json:"leaf"`
	Relevant      bool         `json:"relevant"`
	MkbClass      *MkbClass    `bun:"rel:belongs-to" json:"mkbClass"`
	MkbClassID    uuid.UUID    `bun:"type:uuid" json:"mkbClassId"`
	MkbSubGroups  MkbSubGroups `bun:"rel:has-many" json:"mkbSubGroups"`
	MkbDiagnosis  MkbDiagnoses `bun:"rel:has-many"json:"mkbDiagnosis"`
}

type MkbGroups []*MkbGroup
