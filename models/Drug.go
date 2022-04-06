package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Drug struct {
	bun.BaseModel `bun:"drugs,alias:drugs"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`

	DrugRegimens          []*DrugRegimen `bun:"rel:has-many" json:"drugRegimens"`
	DrugRegimensForDelete []string       `bun:"-" json:"drugRegimensForDelete"`

	DrugsDiagnosis          DrugsDiagnosis `bun:"rel:has-many" json:"drugsDiagnosis"`
	DrugsDiagnosisForDelete []uuid.UUID    `bun:"-" json:"drugsDiagnosisForDelete"`
}

func (item *Drug) SetIdForChildren() {
	if len(item.DrugRegimens) > 0 {
		for i := range item.DrugRegimens {
			item.DrugRegimens[i].DrugID = item.ID
		}
	}
	if len(item.DrugsDiagnosis) > 0 {
		for i := range item.DrugsDiagnosis {
			item.DrugsDiagnosis[i].DrugID = item.ID
		}
	}
}
