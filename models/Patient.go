package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Patient struct {
	bun.BaseModel `bun:"patient,alias:patient"`
	ModelInfo
	ID      uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Human   *Human    `bun:"rel:belongs-to" json:"human"`
	HumanID uuid.UUID `bun:"type:uuid" json:"humanId"`

	RepresentativeToPatient          []*RepresentativeToPatient `bun:"rel:has-many" json:"representativeToPatient"`
	RepresentativeToPatientForDelete []string                   `bun:"-" json:"representativeToPatientForDelete"`
	HeightWeight                     []*HeightWeight            `bun:"rel:has-many" json:"heightWeight"`
	HeightWeightForDelete            []string                   `bun:"-" json:"heightWeightForDelete"`
	Disabilities                     []*Disability              `bun:"rel:has-many" json:"disabilities"`
	DisabilitiesForDelete            []string                   `bun:"-" json:"disabilitiesForDelete"`

	PatientDiagnosis          []*PatientDiagnosis `bun:"rel:has-many" json:"patientDiagnosis"`
	PatientDiagnosisForDelete []string            `bun:"-" json:"patientDiagnosisForDelete"`

	PatientDrugRegimen          []*PatientDrugRegimen `bun:"rel:has-many" json:"patientDrugRegimen"`
	PatientDrugRegimenForDelete []string              `bun:"-" json:"patientDrugRegimenForDelete"`

	RegisterToPatient          []*RegisterToPatient `bun:"rel:has-many" json:"registerToPatient"`
	RegisterToPatientForDelete []string             `bun:"-" json:"registerToPatientForDelete"`

	RegisterPropertyToPatient             []*RegisterPropertyToPatient    `bun:"rel:has-many" json:"registerPropertyToPatient"`
	RegisterPropertySetToPatient          []*RegisterPropertySetToPatient `bun:"rel:has-many" json:"registerPropertySetToPatient"`
	RegisterPropertySetToPatientForDelete []string                        `bun:"-" json:"registerPropertySetToPatientForDelete"`
}

type Patients []*Patient

type PatientsWithCount struct {
	Patients Patients `json:"patients"`
	Count    int      `json:"count"`
}

func (item *Patient) SetFilePath(fileId *string) *string {
	path := item.Human.SetFilePath(fileId)
	if path != nil {
		return path
	}
	for i := range item.Disabilities {
		path := item.Disabilities[i].SetFilePath(fileId)
		if path != nil {
			return path
		}
	}
	return nil
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
	if len(item.PatientDrugRegimen) > 0 {
		for i := range item.PatientDrugRegimen {
			item.PatientDrugRegimen[i].PatientID = item.ID
		}
	}
	if len(item.RegisterToPatient) > 0 {
		for i := range item.RegisterToPatient {
			item.RegisterToPatient[i].PatientID = item.ID
		}
	}
	if len(item.RegisterPropertyToPatient) > 0 {
		for i := range item.RegisterPropertyToPatient {
			item.RegisterPropertyToPatient[i].PatientID = item.ID
		}
	}
	if len(item.RegisterPropertySetToPatient) > 0 {
		for i := range item.RegisterPropertySetToPatient {
			item.RegisterPropertySetToPatient[i].PatientID = item.ID
		}
	}
}

type RepresentativeToPatient struct {
	bun.BaseModel        `bun:"representative_to_patient,alias:representative_to_patient"`
	ID                   uuid.UUID           `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	RepresentativeType   *RepresentativeType `bun:"rel:belongs-to" json:"representativeType"`
	RepresentativeTypeID uuid.UUID           `bun:"type:uuid" json:"representativeTypeId"`
	PatientID            uuid.UUID           `bun:"type:uuid" json:"patientId"`
	Patient              *Patient            `bun:"rel:belongs-to" json:"patient"`
	RepresentativeID     uuid.UUID           `bun:"type:uuid" json:"representativeId"`
	Representative       *Representative     `bun:"rel:belongs-to" json:"representative"`
}

type RegisterPropertySetToPatient struct {
	bun.BaseModel         `bun:"register_property_set_to_patient,alias:register_property_set_to_patient"`
	ID                    uuid.UUID            `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	RegisterPropertySet   *RegisterPropertySet `bun:"rel:belongs-to" json:"registerPropertySet"`
	RegisterPropertySetID uuid.UUID            `bun:"type:uuid" json:"registerPropertySetId"`
	Patient               *Patient             `bun:"rel:has-one" json:"patient"`
	PatientID             uuid.UUID            `bun:"type:uuid" json:"patientId"`
}

type RegisterToPatient struct {
	bun.BaseModel `bun:"register_to_patient,alias:register_to_patient"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Register      *Register `bun:"rel:belongs-to" json:"register"`
	RegisterID    uuid.UUID `bun:"type:uuid" json:"registerId"`
	Patient       *Patient  `bun:"rel:belongs-to" json:"patient"`
	PatientID     uuid.UUID `bun:"type:uuid" json:"patientId"`
}

type RegisterPropertyToPatient struct {
	bun.BaseModel `bun:"register_property_to_patient,alias:register_property_to_patient"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `

	ValueString string    `json:"valueString"`
	ValueNumber int       `json:"valueNumber"`
	ValueDate   time.Time `json:"valueDate"`
	ValueOther  string    `json:"valueOther"`

	RegisterPropertyRadio   *RegisterPropertyRadio `bun:"rel:belongs-to" json:"registerPropertyRadio"`
	RegisterPropertyRadioID uuid.NullUUID          `bun:"type:uuid" json:"registerPropertyRadioId"`

	RegisterProperty   *RegisterProperty `bun:"rel:belongs-to" json:"registerProperty"`
	RegisterPropertyID uuid.UUID         `bun:"type:uuid" json:"registerPropertyId"`
	Patient            *Patient          `bun:"rel:has-one" json:"patient"`
	PatientID          uuid.UUID         `bun:"type:uuid" json:"patientId"`
}
