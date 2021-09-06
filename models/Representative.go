package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Representative struct {
	bun.BaseModel `bun:"representative,alias:representative"`
	ModelInfo
	ID uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `

	Human   *Human    `bun:"rel:belongs-to" json:"human"`
	HumanID uuid.UUID `bun:"type:uuid" json:"humanId"`

	RepresentativeToPatient          []*RepresentativeToPatient `bun:"rel:has-many" json:"representativeToPatient"`
	RepresentativeToPatientForDelete []string                   `bun:"-" json:"representativeToPatientForDelete"`
}

func (item *Representative) SetFilePath(fileId *string) *string {
	path := item.Human.SetFilePath(fileId)
	if path != nil {
		return path
	}
	return nil
}

func (item *Representative) SetIdForChildren() {
	if len(item.RepresentativeToPatient) > 0 {
		for i := range item.RepresentativeToPatient {
			item.RepresentativeToPatient[i].RepresentativeID = item.ID
		}
	}
}
