package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Patient struct {
	bun.BaseModel `bun:"patients,alias:patients"`
	ModelInfo
	ID      uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Human   *Human    `bun:"rel:belongs-to" json:"human"`
	HumanID uuid.UUID `bun:"type:uuid" json:"humanId"`

	Region   *Region   `bun:"rel:belongs-to" json:"region"`
	RegionID uuid.UUID `bun:"type:uuid" json:"regionId"`

	RepresentativeToPatient          []*RepresentativeToPatient `bun:"rel:has-many" json:"representativeToPatient"`
	RepresentativeToPatientForDelete []uuid.UUID                `bun:"-" json:"representativeToPatientForDelete"`
	HeightWeight                     []*HeightWeight            `bun:"rel:has-many" json:"heightWeight"`
	HeightWeightForDelete            []uuid.UUID                `bun:"-" json:"heightWeightForDelete"`
	Disabilities                     []*Disability              `bun:"rel:has-many" json:"disabilities"`
	DisabilitiesForDelete            []uuid.UUID                `bun:"-" json:"disabilitiesForDelete"`
	ChestCircumference               []*ChestCircumference      `bun:"rel:has-many" json:"chestCircumference"`
	ChestCircumferenceForDelete      []uuid.UUID                `bun:"-" json:"chestCircumferenceForDelete"`
	HeadCircumference                []*HeadCircumference       `bun:"rel:has-many" json:"headCircumference"`
	HeadCircumferenceForDelete       []uuid.UUID                `bun:"-" json:"headCircumferenceForDelete"`

	PatientDiagnosis          []*PatientDiagnosis `bun:"rel:has-many" json:"patientDiagnosis"`
	PatientDiagnosisForDelete []uuid.UUID         `bun:"-" json:"patientDiagnosisForDelete"`

	PatientDrugRegimen          []*PatientDrugRegimen `bun:"rel:has-many" json:"patientDrugRegimen"`
	PatientDrugRegimenForDelete []uuid.UUID           `bun:"-" json:"patientDrugRegimenForDelete"`

	RegisterToPatient          []*RegisterToPatient `bun:"rel:has-many" json:"registerToPatient"`
	RegisterToPatientForDelete []uuid.UUID          `bun:"-" json:"registerToPatientForDelete"`

	RegisterPropertyToPatient             []*RegisterPropertyToPatient    `bun:"rel:has-many" json:"registerPropertyToPatient"`
	RegisterPropertyOthersPatient         RegisterPropertyOthersToPatient `bun:"rel:has-many" json:"registerPropertyOthersToPatient"`
	RegisterPropertyToPatientForDelete    []uuid.UUID                     `bun:"-" json:"registerPropertyToPatientForDelete"`
	RegisterPropertySetToPatient          []*RegisterPropertySetToPatient `bun:"rel:has-many" json:"registerPropertySetToPatient"`
	RegisterPropertySetToPatientForDelete []uuid.UUID                     `bun:"-" json:"registerPropertySetToPatientForDelete"`
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
	if len(item.ChestCircumference) > 0 {
		for i := range item.ChestCircumference {
			item.ChestCircumference[i].PatientID = item.ID
		}
	}
	if len(item.HeadCircumference) > 0 {
		for i := range item.HeadCircumference {
			item.HeadCircumference[i].PatientID = item.ID
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
	if len(item.RegisterPropertyOthersPatient) > 0 {
		for i := range item.RegisterPropertyOthersPatient {
			item.RegisterPropertyOthersPatient[i].PatientID = item.ID
		}
	}
}

func (item *Patient) SetDeleteIdForChildren() {
	for i := range item.RepresentativeToPatient {
		item.RepresentativeToPatientForDelete = append(item.RepresentativeToPatientForDelete, item.RepresentativeToPatient[i].ID)
	}
	for i := range item.HeightWeight {
		item.HeightWeightForDelete = append(item.HeightWeightForDelete, item.HeightWeight[i].ID)
	}
	for i := range item.ChestCircumference {
		item.ChestCircumferenceForDelete = append(item.ChestCircumferenceForDelete, item.ChestCircumference[i].ID)
	}
	for i := range item.HeadCircumference {
		item.HeadCircumferenceForDelete = append(item.HeadCircumferenceForDelete, item.HeadCircumference[i].ID)
	}
	for i := range item.Disabilities {
		item.DisabilitiesForDelete = append(item.DisabilitiesForDelete, item.Disabilities[i].ID)
	}
	for i := range item.PatientDiagnosis {
		item.PatientDiagnosisForDelete = append(item.PatientDiagnosisForDelete, item.PatientDiagnosis[i].ID)
	}
	for i := range item.PatientDrugRegimen {
		item.PatientDrugRegimenForDelete = append(item.PatientDrugRegimenForDelete, item.PatientDrugRegimen[i].ID)
	}
	for i := range item.RegisterToPatient {
		item.RegisterToPatientForDelete = append(item.RegisterToPatientForDelete, item.RegisterToPatient[i].ID)
	}
	for i := range item.RegisterPropertyToPatient {
		item.RegisterPropertyToPatientForDelete = append(item.RegisterPropertyToPatientForDelete, item.RegisterPropertyToPatient[i].ID)
	}
	for i := range item.RegisterPropertySetToPatient {
		item.RegisterPropertySetToPatientForDelete = append(item.RegisterPropertySetToPatientForDelete, item.RegisterPropertySetToPatient[i].ID)
	}
}
