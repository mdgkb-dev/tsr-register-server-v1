package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Patient struct {
	bun.BaseModel `bun:"patient,alias:patient"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Human         *Human    `bun:"rel:belongs-to" json:"human"`
	HumanID       uuid.UUID `bun:"type:uuid" json:"humanId"`

	RepresentativeToPatient          []*RepresentativeToPatient `bun:"rel:has-many" json:"representativeToPatient"`
	RepresentativeToPatientForDelete []string                   `bun:"-" json:"representativeToPatientForDelete"`
	AnthropometryData                []*AnthropometryData       `bun:"rel:has-many" json:"anthropometryData"`
	AnthropometryDataForDelete       []string                   `bun:"-" json:"anthropometryDataForDelete"`
	Disabilities                     []*Disability              `bun:"rel:has-many" json:"disabilities"`

	PatientDiagnosis []*PatientDiagnosis `bun:"rel:has-many" json:"PatientDiagnosis"`
}

type RepresentativeToPatient struct {
	bun.BaseModel        `bun:"representative_to_patient,alias:representative_to_patient"`
	ID                   uuid.UUID           `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	RepresentativeType   *RepresentativeType `bun:"rel:belongs-to" json:"representativeType"`
	RepresentativeTypeID uuid.UUID           `bun:"type:uuid" json:"representativeTypeId"`
	PatientID            uuid.UUID           `bun:"type:uuid" json:"PatientId"`
	Patient              *Patient            `bun:"rel:belongs-to" json:"patient"`
	RepresentativeID     uuid.UUID           `bun:"type:uuid" json:"representativeId"`
	Representative       *Representative     `bun:"rel:belongs-to" json:"representative"`
}

type PatientDiagnosis struct {
	bun.BaseModel     `bun:"patient_diagnosis,alias:patient_diagnosis"`
	ID                uuid.UUID        `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Patient           *Patient         `bun:"rel:has-one" json:"patient"`
	PatientID         uuid.UUID        `bun:"type:uuid" json:"PatientId"`
	MkbDiagnosis      *MkbDiagnosis    `bun:"rel:belongs-to" json:"mkbDiagnosis"`
	MkbDiagnosisID    uuid.UUID        `bun:"type:uuid" json:"mkbDiagnosisId"`
	MkbSubDiagnosis   *MkbSubDiagnosis `bun:"rel:has-one" json:"mkbSubDiagnosis"`
	MkbSubDiagnosisID uuid.UUID        `bun:"type:uuid" json:"mkbSubDiagnosisId"`
	Primary           bool             `json:"primary"`
}

type PatientDiagnosisAnamnesis struct {
	bun.BaseModel      `bun:"patient_diagnosis_anamnesis,alias:patient_diagnosis_anamnesis"`
	ID                 uuid.UUID         `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	PatientDiagnosis   *PatientDiagnosis `bun:"rel:belongs-to" json:"patientDiagnosis"`
	PatientDiagnosisID uuid.UUID         `bun:"type:uuid" json:"patientDiagnosisId"`
	Value              string            `json:"value"`
	Date               time.Time         `json:"date"`
}

type RegisterPropertySetToPatient struct {
	bun.BaseModel         `bun:"register_property_set_to_patient,alias:register_property_set_to_patient"`
	ID                    uuid.UUID            `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	RegisterPropertySet   *RegisterPropertySet `bun:"rel:belongs-to" json:"registerPropertySet"`
	RegisterPropertySetID uuid.UUID            `bun:"type:uuid" json:"registerPropertySetId"`
	Patient               *Patient             `bun:"rel:has-one" json:"patient"`
	PatientID             uuid.UUID            `bun:"type:uuid" json:"PatientId"`
}

type RegisterToPatient struct {
	bun.BaseModel `bun:"register_to_patient,alias:register_to_patient"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Register      *Register `bun:"rel:belongs-to" json:"register"`
	RegisterID    uuid.UUID `bun:"type:uuid" json:"registerId"`
	Patient       *Patient  `bun:"rel:has-one" json:"patient"`
	PatientID     uuid.UUID `bun:"type:uuid" json:"PatientId"`
}

type RegisterPropertyToPatient struct {
	bun.BaseModel `bun:"register_property_to_patient,alias:register_property_to_patient"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `

	ValueString string `json:"valueString"`
	ValueNumber string `json:"valueNumber"`
	ValueDate   string `json:"valueDate"`
	ValueOther  string `json:"valueOther"`

	RegisterPropertyRadio   *RegisterPropertyRadio `bun:"rel:belongs-to" json:"registerPropertyRadio"`
	RegisterPropertyRadioID uuid.UUID              `bun:"type:uuid" json:"registerPropertyRadioId"`

	RegisterProperty   *RegisterProperty `bun:"rel:belongs-to" json:"registerProperty"`
	RegisterPropertyID uuid.UUID         `bun:"type:uuid" json:"registerPropertyId"`
	Patient            *Patient          `bun:"rel:has-one" json:"patient"`
	PatientID          uuid.UUID         `bun:"type:uuid" json:"PatientId"`
}
