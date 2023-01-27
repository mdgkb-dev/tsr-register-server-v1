package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Register struct {
	bun.BaseModel              `bun:"register,alias:register"`
	ID                         uuid.UUID            `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                       string               `json:"name"`
	RegisterGroups             RegisterGroups       `bun:"rel:has-many" json:"registerGroups"`
	RegisterGroupsForDelete    []uuid.UUID          `bun:"-" json:"registerGroupsForDelete"`
	RegisterDiagnosis          []*RegisterDiagnosis `bun:"rel:has-many" json:"registerDiagnosis"`
	RegisterDiagnosisForDelete []string             `bun:"-" json:"registerDiagnosisForDelete"`

	RegisterToPatient      []*RegisterToPatient `bun:"rel:has-many" json:"registerToPatient"`
	RegisterToPatientCount int                  `bun:"-" json:"registerToPatientCount"`
}

func (item *Register) SetIDForChildren() {
	if len(item.RegisterGroups) > 0 {
		for i := range item.RegisterGroups {
			item.RegisterGroups[i].RegisterID = item.ID
		}
	}
	if len(item.RegisterDiagnosis) > 0 {
		for i := range item.RegisterDiagnosis {
			item.RegisterDiagnosis[i].RegisterID = item.ID
		}
	}
}

func (item *Register) GetPatientsAverageAge() int {
	sum := 0
	for _, p := range item.RegisterToPatient {
		sum += p.Patient.Human.GetAge()
	}
	res := sum / len(item.RegisterToPatient)
	return res
}

type RegisterDiagnosis struct {
	bun.BaseModel          `bun:"register_diagnosis,alias:register_diagnosis"`
	ID                     uuid.UUID             `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Register               *Register             `bun:"rel:belongs-to" json:"register"`
	RegisterID             uuid.UUID             `bun:"type:uuid" json:"registerId"`
	MkbDiagnosis           *MkbDiagnosis         `bun:"rel:belongs-to" json:"mkbDiagnosis"`
	MkbDiagnosisID         uuid.UUID             `bun:"type:uuid" json:"mkbDiagnosisId"`
	MkbSubDiagnosis        *MkbDiagnosis         `bun:"rel:belongs-to" json:"mkbSubDiagnosis"`
	MkbSubDiagnosisID      uuid.NullUUID         `bun:"type:uuid" json:"mkbSubDiagnosisId"`
	MkbConcreteDiagnosis   *MkbConcreteDiagnosis `bun:"rel:belongs-to" json:"mkbConcreteDiagnosis"`
	MkbConcreteDiagnosisID uuid.NullUUID         `bun:"type:uuid,nullzero" json:"mkbConcreteDiagnosisId"`
}
