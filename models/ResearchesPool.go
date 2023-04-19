package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResearchesPool struct {
	bun.BaseModel                      `bun:"researches_pools,alias:researches_pools"`
	ID                                 uuid.NullUUID             `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                               string                    `json:"name"`
	ResearchesPoolsResearches          ResearchesPoolsResearches `bun:"rel:has-many" json:"researchesPoolsResearches"`
	ResearchesPoolsResearchesForDelete []uuid.UUID               `bun:"-" json:"researchesPoolsResearchesForDelete"`

	PatientsResearchesPools          PatientsResearchesPools `bun:"rel:has-many" json:"patientsResearchesPools"`
	PatientsResearchesPoolsForDelete []uuid.UUID             `bun:"-" json:"patientsResearchesPoolsForDelete"`
	//ResearchDiagnosis          []*ResearchDiagnosis `bun:"rel:has-many" json:"ResearchDiagnosis"`
	//ResearchDiagnosisForDelete []string             `bun:"-" json:"ResearchDiagnosisForDelete"`

	//PatientsResearches []*ResearchResults `bun:"rel:has-many" json:"patientsResearches"`
	//ResearchToPatientCount int                   `bun:"-" json:"ResearchToPatientCount"`
}

type ResearchesPools []*ResearchesPool
type ResearchesPoolsWithCount struct {
	ResearchesPools ResearchesPools `json:"items"`
	Count           int             `json:"count"`
}

func (item *ResearchesPool) SetIDForChildren() {
	//if len(item.Researches) > 0 {
	//	for i := range item.Researches {
	//		item.Researches[i].ResearchPoolID = item.ID
	//	}
	//}
	//if len(item.ResearchDiagnosis) > 0 {
	//	for i := range item.ResearchDiagnosis {
	//		item.ResearchDiagnosis[i].ResearchPoolID = item.ID
	//	}
	//}
}

//func (item *ResearchesPool) GetPatientsAverageAge() int {
//	sum := 0
//	for _, p := range item.ResearchResults {
//		sum += p.Patient.Human.GetAge()
//	}
//	res := sum / len(item.ResearchResults)
//	return res
//}
