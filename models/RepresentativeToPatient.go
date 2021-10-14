package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RepresentativeToPatient struct {
	bun.BaseModel        `bun:"representative_to_patient,alias:representative_to_patient"`
	ID                   uuid.UUID           `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	RepresentativeType   *RepresentativeType `bun:"rel:belongs-to" json:"representativeType"`
	RepresentativeTypeID uuid.UUID           `bun:"type:uuid" json:"representativeTypeId"`
	PatientID            uuid.UUID           `bun:"type:uuid" json:"patientId"`
	Patient              *Patient            `bun:"rel:belongs-to" json:"patient"`
	RepresentativeID     uuid.UUID           `bun:"type:uuid" json:"representativeId"`
	Representative       *Representative     `bun:"rel:belongs-to" json:"representative"`
	DeletedAt            time.Time           `bun:",soft_delete" json:"deletedAt"`
}
