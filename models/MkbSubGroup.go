package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MkbSubGroup struct {
	bun.BaseModel   `bun:"mkb_sub_group,alias:mkb_sub_group"`
	ID              uuid.UUID         `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name            string            `json:"name"`
	RangeStart      string            `json:"rangeStart"`
	RangeEnd        string            `json:"rangeEnd"`
	Comment         string            `json:"comment"`
	Leaf            bool              `json:"leaf"`
	Relevant        bool              `json:"relevant"`
	MkbGroup        *MkbGroup         `bun:"rel:belongs-to" json:"mkbGroup"`
	MkbGroupID      uuid.UUID         `bun:"type:uuid" json:"mkbGroupId"`
	MkbSubSubGroups []*MkbSubSubGroup `bun:"rel:has-many" json:"mkbSubSubGroups"`
	MkbDiagnosis    []*MkbDiagnosis   `bun:"rel:has-many"json:"mkbDiagnosis"`
}

type MkbSubGroups []*MkbSubGroup
