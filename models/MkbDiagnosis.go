package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MkbDiagnosis struct {
	bun.BaseModel    `bun:"mkb_diagnosis,alias:mkb_diagnosis"`
	ID               uuid.UUID          `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name             string             `json:"name"`
	Code             string             `json:"code"`
	Comment          string             `json:"comment"`
	Leaf             bool               `json:"bool"`
	Relevant         bool               `json:"relevant"`
	MkbClass         *MkbClass          `bun:"rel:belongs-to" json:"mkbClass"`
	MkbClassID       uuid.UUID          `bun:"type:uuid" json:"mkbClassId"`
	MkbGroup         *MkbGroup          `bun:"rel:belongs-to" json:"mkbGroup"`
	MkbGroupID       uuid.UUID          `bun:"type:uuid" json:"mkbGroupId"`
	MkbSubGroup      *MkbSubGroup       `bun:"rel:belongs-to" json:"mkbSubGroup"`
	MkbSubGroupID    uuid.UUID          `bun:"type:uuid" json:"mkbSubGroupId"`
	MkbSubSubGroup   *MkbSubSubGroup    `bun:"rel:belongs-to" json:"mkbSubSubGroup"`
	MkbSubSubGroupID uuid.UUID          `bun:"type:uuid" json:"mkbSubSubGroupId"`
	MkbSubDiagnosis  []*MkbSubDiagnosis `bun:"rel:has-many" json:"mkbSubDiagnosis"`
}
