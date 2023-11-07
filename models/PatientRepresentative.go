package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PatientRepresentative struct {
	bun.BaseModel        `bun:"patients_representatives,alias:patients_representatives"`
	ID                   uuid.NullUUID           `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	RepresentativeType   *RepresentativeType `bun:"rel:belongs-to" json:"representativeType"`
	RepresentativeTypeID uuid.NullUUID       `bun:"type:uuid" json:"representativeTypeId"`
	PatientID            uuid.NullUUID       `bun:"type:uuid" json:"patientId"`
	Patient              *Patient            `bun:"rel:belongs-to" json:"patient"`
	RepresentativeID     uuid.NullUUID           `bun:"type:uuid" json:"representativeId"`
	Representative       *Representative     `bun:"rel:belongs-to" json:"representative"`
	DeletedAt            *time.Time          `bun:",soft_delete" json:"deletedAt"`
}

type PatientsRepresentatives []*PatientRepresentative

type PatientsRepresentativesWithCount struct {
	PatientsRepresentatives PatientsRepresentatives `json:"items"`
	Count                   int                     `json:"count"`
}
