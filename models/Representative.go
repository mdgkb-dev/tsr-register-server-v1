package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Representative struct {
	bun.BaseModel `bun:"representatives,select:representatives_view,alias:representatives_view"`
	ModelInfo
	ID uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Human   *Human    `bun:"rel:belongs-to" json:"human"`
	HumanID uuid.UUID `bun:"type:uuid" json:"humanId"`

	PatientsRepresentatives          PatientsRepresentatives `bun:"rel:has-many" json:"patientsRepresentatives"`
	PatientsRepresentativesForDelete []uuid.UUID             `bun:"-" json:"patientsRepresentativesForDelete"`
	FullName                         string                  `bun:"-" json:"fullName"`
	IsMale                           string                  `bun:"-" json:"isMale"`
	DateBirth                        string                  `bun:"-" json:"dateBirth"`
}

type Representatives []*Representative

type RepresentativesWithCount struct {
	Representatives Representatives `json:"items"`
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
	if len(item.PatientsRepresentatives) > 0 {
		for i := range item.PatientsRepresentatives {
			item.PatientsRepresentatives[i].RepresentativeID = item.ID
		}
	}
}
