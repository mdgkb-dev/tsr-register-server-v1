package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MkbClass struct {
	bun.BaseModel `bun:"mkb_class,alias:mkb_class"`
	ID            uuid.UUID    `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string       `json:"name"`
	Number        string       `json:"number"`
	RangeStart    string       `json:"rangeStart"`
	RangeEnd      string       `json:"rangeEnd"`
	Comment       string       `json:"comment"`
	Leaf          bool         `json:"leaf"`
	Relevant      bool         `json:"relevant"`
	MkbGroups     MkbGroups    `bun:"rel:has-many" json:"mkbGroups"`
	MkbDiagnosis  MkbDiagnoses `bun:"rel:has-many" json:"mkbDiagnosis"`
}

type MkbClasses []*MkbClass
