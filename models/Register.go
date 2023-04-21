package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Register struct {
	bun.BaseModel           `bun:"register,alias:register"`
	ID                      uuid.UUID   `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                    string      `json:"name"`
	RegisterGroups          Researches  `bun:"rel:has-many" json:"registerGroups"`
	RegisterGroupsForDelete []uuid.UUID `bun:"-" json:"registerGroupsForDelete"`
	//RegisterDiagnosis          []*RegisterDiagnosis `bun:"rel:has-many" json:"registerDiagnosis"`
	RegisterDiagnosisForDelete []string `bun:"-" json:"registerDiagnosisForDelete"`

	RegisterToPatient      []*ResearchResult `bun:"rel:has-many" json:"registerToPatient"`
	RegisterToPatientCount int               `bun:"-" json:"registerToPatientCount"`
}

func (item *Register) SetIDForChildren() {
	//if len(item.RegisterGroups) > 0 {
	//	for i := range item.RegisterGroups {
	//		item.RegisterGroups[i].ResearchPoolID = item.ID
	//	}
	//}
	//if len(item.RegisterDiagnosis) > 0 {
	//	for i := range item.RegisterDiagnosis {
	//		item.RegisterDiagnosis[i].ResearchPoolID = item.ID
	//	}
	//}
}

func (item *Register) GetPatientsAverageAge() int {
	sum := 0
	//for _, _ := range item.RegisterToPatient {
	//	sum += p.Patient.Human.GetAge()
	//}
	res := sum / len(item.RegisterToPatient)
	return res
}
