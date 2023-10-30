package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Drug struct {
	bun.BaseModel `bun:"drugs,select:drugs_view,alias:drugs_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	NameINN       string        `bun:"name_inn" json:"nameINN"`
	DrugRegimens  DrugRegimens  `bun:"rel:has-many" json:"drugRegimens"`
	DrugForms     DrugForms     `bun:"rel:has-many" json:"drugForms"`

	DrugsDiagnosis DrugsDiagnosis `bun:"rel:has-many" json:"drugsDiagnosis"`

	DrugsDiagnosisForDelete []uuid.UUID `bun:"-" json:"drugsDiagnosisForDelete"`
}

type Drugs []*Drug
type DrugsWithCount struct {
	Drugs Drugs `json:"items"`
	Count int   `json:"count"`
}

func (item *Drug) SetIDForChildren() {
	if len(item.DrugsDiagnosis) > 0 {
		for i := range item.DrugsDiagnosis {
			item.DrugsDiagnosis[i].DrugID = item.ID
		}
	}
}


