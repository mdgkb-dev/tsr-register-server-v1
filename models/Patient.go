package models

import (
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Patient struct {
	bun.BaseModel `bun:"patient,alias:patient"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Human         *Human    `bun:"rel:belongs-to" json:"human"`
	HumanID       uuid.UUID `bun:"type:uuid" json:"humanId"`

	RepresentativeToPatient          []*RepresentativeToPatient `bun:"rel:has-many" json:"representativeToPatient"`
	RepresentativeToPatientForDelete []string                   `bun:"-" json:"representativeToPatientForDelete"`
	HeightWeight                     []*HeightWeight            `bun:"rel:has-many" json:"heightWeight"`
	HeightWeightForDelete            []string                   `bun:"-" json:"heightWeightForDelete"`
	Disabilities                     []*Disability              `bun:"rel:has-many" json:"disabilities"`
	DisabilitiesForDelete            []string                   `bun:"-" json:"disabilitiesForDelete"`

	PatientDiagnosis          []*PatientDiagnosis `bun:"rel:has-many" json:"patientDiagnosis"`
	PatientDiagnosisForDelete []string            `bun:"-" json:"patientDiagnosisForDelete"`

	RegisterToPatient          []*RegisterToPatient `bun:"rel:has-many" json:"registerToPatient"`
	RegisterToPatientForDelete []string             `bun:"-" json:"registerToPatientForDelete"`

	RegisterPropertyToPatient    []*RegisterPropertyToPatient    `bun:"rel:has-many" json:"registerPropertyToPatient"`
	RegisterPropertySetToPatient []*RegisterPropertySetToPatient `bun:"rel:has-many" json:"registerPropertySetToPatient"`
}

func (item *Patient) SetFilePath(parentId *string) *string {
	newPath := filepath.Join(item.Human.ID.String())

	for i := range item.Human.Documents {
		path := item.Human.Documents[i].SetFilePath(parentId, &newPath)
		if path != "" {
			newPath = filepath.Join(newPath, path)
		}
	}
	for i := range item.Disabilities {
		item.Disabilities[i].SetFilePath(parentId, &newPath)
	}
	return &newPath
}

func (item *Patient) SetIdForChildren() {
	if len(item.RepresentativeToPatient) > 0 {
		for i := range item.RepresentativeToPatient {
			item.RepresentativeToPatient[i].PatientID = item.ID
		}
	}
	if len(item.HeightWeight) > 0 {
		for i := range item.HeightWeight {
			item.HeightWeight[i].PatientID = item.ID
		}
	}
	if len(item.Disabilities) > 0 {
		for i := range item.Disabilities {
			item.Disabilities[i].PatientID = item.ID
		}
	}
	if len(item.PatientDiagnosis) > 0 {
		for i := range item.PatientDiagnosis {
			item.PatientDiagnosis[i].PatientID = item.ID
		}
	}
	if len(item.RegisterToPatient) > 0 {
		for i := range item.RegisterToPatient {
			item.RegisterToPatient[i].PatientID = item.ID
		}
	}
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
	MkbSubDiagnosis   *MkbSubDiagnosis `bun:"rel:belongs-to" json:"mkbSubDiagnosis"`
	MkbSubDiagnosisID uuid.NullUUID    `bun:"type:uuid,nullzero" json:"mkbSubDiagnosisId"`
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
	Patient       *Patient  `bun:"rel:belongs-to" json:"patient"`
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
