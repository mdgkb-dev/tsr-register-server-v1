package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Representative struct {
	bun.BaseModel `bun:"representative,select:representatives_view,alias:representatives_view"`
	ModelInfo
	ID uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Human   *Human    `bun:"rel:belongs-to" json:"human"`
	HumanID uuid.UUID `bun:"type:uuid" json:"humanId"`

	RepresentativeToPatient          []*RepresentativeToPatient `bun:"rel:has-many" json:"representativeToPatient"`
	RepresentativeToPatientForDelete []uuid.UUID                `bun:"-" json:"representativeToPatientForDelete"`
}

type Representatives []*Representative

type RepresentativesWithCount struct {
	Representatives Representatives `json:"representatives"`
	Count           int             `json:"count"`
}

func (item *Representative) SetFilePath(fileID *string) *string {
	path := item.Human.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (item *Representative) SetIDForChildren() {
	if len(item.RepresentativeToPatient) > 0 {
		for i := range item.RepresentativeToPatient {
			item.RepresentativeToPatient[i].RepresentativeID = item.ID
		}
	}
}
