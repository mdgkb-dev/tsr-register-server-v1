package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Register struct {
	bun.BaseModel                    `bun:"register,alias:register"`
	ID                               uuid.UUID                  `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                             string                     `json:"name"`
	RegisterGroupToRegister          []*RegisterGroupToRegister `bun:"rel:has-many" json:"registerGroupToRegister"`
	RegisterGroupToRegisterForDelete []string                   `bun:"-" json:"registerGroupToRegisterForDelete"`
	RegisterDiagnosis                []*RegisterDiagnosis       `bun:"rel:has-many" json:"registerDiagnosis"`
	RegisterDiagnosisForDelete       []string                   `bun:"-" json:"registerDiagnosisForDelete"`
}

func (item *Register) SetIdForChildren() {
	if len(item.RegisterGroupToRegister) > 0 {
		for i := range item.RegisterGroupToRegister {
			item.RegisterGroupToRegister[i].RegisterID = item.ID
		}
	}
	if len(item.RegisterDiagnosis) > 0 {
		for i := range item.RegisterDiagnosis {
			item.RegisterDiagnosis[i].RegisterID = item.ID
		}
	}
}

type RegisterGroupToRegister struct {
	bun.BaseModel   `bun:"register_group_to_register,alias:register_group_to_register"`
	ID              uuid.UUID      `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	RegisterID      uuid.UUID      `bun:"type:uuid" json:"registerId"`
	Register        *Register      `bun:"rel:has-one" json:"register"`
	RegisterGroupID uuid.UUID      `bun:"type:uuid" json:"registerGroupId"`
	RegisterGroup   *RegisterGroup `bun:"rel:has-one" json:"registerGroup"`
}

type RegisterDiagnosis struct {
	bun.BaseModel     `bun:"register_diagnosis,alias:register_diagnosis"`
	ID                uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	RegisterID        uuid.UUID     `bun:"type:uuid" json:"registerId"`
	Register          *Register     `bun:"rel:has-one" json:"register"`
	MkbDiagnosisID    uuid.UUID     `bun:"type:uuid" json:"mkbDiagnosisId"`
	MkbDiagnosis      *MkbDiagnosis `bun:"rel:has-one" json:"mkbDiagnosis"`
	MkbSubDiagnosisID uuid.UUID     `bun:"type:uuid" json:"mkbSubDiagnosisId"`
	MkbSubDiagnosis   *MkbDiagnosis `bun:"rel:has-one" json:"mkbSubDiagnosis"`
}
