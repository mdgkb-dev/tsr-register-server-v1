package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Patient struct {
	bun.BaseModel `bun:"patients,select:patients_view,alias:patients_view"`
	ModelInfo
	ID      uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Human   *Human        `bun:"rel:belongs-to" json:"human"`
	HumanID uuid.UUID     `bun:"type:uuid" json:"humanId"`

	Region   *Region   `bun:"rel:belongs-to" json:"region"`
	RegionID uuid.UUID `bun:"type:uuid" json:"regionId"`

	PatientsRepresentatives          PatientsRepresentatives `bun:"rel:has-many" json:"patientsRepresentatives"`
	PatientsRepresentativesForDelete []uuid.UUID             `bun:"-" json:"patientsRepresentativesForDelete"`
	HeightWeight                     []*HeightWeight         `bun:"rel:has-many" json:"heightWeight"`
	HeightWeightForDelete            []uuid.UUID             `bun:"-" json:"heightWeightForDelete"`
	Disabilities                     Disabilities            `bun:"rel:has-many" json:"disabilities"`
	DisabilitiesForDelete            []uuid.UUID             `bun:"-" json:"disabilitiesForDelete"`
	ChestCircumference               []*ChestCircumference   `bun:"rel:has-many" json:"chestCircumference"`
	ChestCircumferenceForDelete      []uuid.UUID             `bun:"-" json:"chestCircumferenceForDelete"`
	HeadCircumference                []*HeadCircumference    `bun:"rel:has-many" json:"headCircumference"`
	HeadCircumferenceForDelete       []uuid.UUID             `bun:"-" json:"headCircumferenceForDelete"`

	PatientDiagnosis          []*PatientDiagnosis `bun:"rel:has-many" json:"patientDiagnosis"`
	PatientDiagnosisForDelete []uuid.UUID         `bun:"-" json:"patientDiagnosisForDelete"`

	PatientDrugRegimen          []*PatientDrugRegimen `bun:"rel:has-many" json:"patientDrugRegimen"`
	PatientDrugRegimenForDelete []uuid.UUID           `bun:"-" json:"patientDrugRegimenForDelete"`

	//RegisterGroupsToPatient           RegisterGroupsToPatients `bun:"rel:has-many" json:"registerGroupsToPatient"`
	//RegisterGroupsToPatientsForDelete []uuid.UUID              `bun:"-" json:"registerGroupsToPatientForDelete"`

	PatientsRegisters PatientsRegisters `bun:"rel:has-many" json:"patientsRegisters"`
	Commissions       Commissions       `bun:"rel:has-many" json:"commissions"`
	//RegisterToPatientForDelete []uuid.UUID       `bun:"-" json:"registerToPatientForDelete"`

	ChopScaleTests          ChopScaleTests `bun:"rel:has-many" json:"chopScaleTests"`
	ChopScaleTestsForDelete []uuid.UUID    `bun:"-" json:"chopScaleTestsForDelete"`

	HmfseScaleTests          HmfseScaleTests `bun:"rel:has-many" json:"hmfseScaleTests"`
	HmfseScaleTestsForDelete []uuid.UUID     `bun:"-" json:"chohmfseScaleTestsForDelete"`

	PatientsResearchesPools          PatientsResearchesPools `bun:"rel:has-many" json:"patientsResearchesPools"`
	PatientsResearchesPoolsForDelete []uuid.UUID             `bun:"-" json:"patientsResearchesPoolsForDelete"`

	PatientsResearches          PatientsResearches `bun:"rel:has-many" json:"patientsResearches"`
	PatientsResearchesForDelete []uuid.UUID        `bun:"-" json:"patientsResearchesForDelete"`
	PatientHistories            PatientHistories   `bun:"rel:has-many" json:"patientHistories"`
	FullName                    string             `bun:"-" json:"fullName"`
	IsMale                      string             `bun:"-" json:"isMale"`
	DateBirth                   string             `bun:"-" json:"dateBirth"`

	CreatedBy   *User         `bun:"rel:belongs-to" json:"createdBy"`
	CreatedByID uuid.NullUUID `bun:"type:uuid" json:"createdById"`
}

type Patients []*Patient

type PatientsWithCount struct {
	Patients Patients `json:"items"`
	Count    int      `json:"count"`
}

func (item *Patient) SetFilePath(fileID *string) *string {
	path := item.Human.SetFilePath(fileID)
	if path != nil {
		return path
	}
	for i := range item.Disabilities {
		path := item.Disabilities[i].SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	//path = item.RegisterGroupsToPatient.SetFilePath(fileID)
	//if path != nil {
	//	return path
	//}
	return nil
}

func (item *Patient) SetIDForChildren() {
	if len(item.PatientsRepresentatives) > 0 {
		for i := range item.PatientsRepresentatives {
			item.PatientsRepresentatives[i].PatientID = item.ID
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
	if len(item.PatientsRegisters) > 0 {
		//for i := range item.RegisterToPatient {
		//	item.RegisterToPatient[i].PatientID = item.ID
		//}
	}
	//if len(item.RegisterGroupsToPatient) > 0 {
	//	for i := range item.RegisterGroupsToPatient {
	//		item.RegisterGroupsToPatient[i].PatientID = item.ID
	//	}
	//}
	if len(item.ChopScaleTests) > 0 {
		for i := range item.ChopScaleTests {
			item.ChopScaleTests[i].PatientID = item.ID
		}
	}
	if len(item.HmfseScaleTests) > 0 {
		for i := range item.HmfseScaleTests {
			item.HmfseScaleTests[i].PatientID = item.ID
		}
	}
}

func (item *Patient) SetDeleteIDForChildren() {
	for i := range item.PatientsRepresentatives {
		item.PatientsRepresentativesForDelete = append(item.PatientsRepresentativesForDelete, item.PatientsRepresentatives[i].ID)
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
	//for i := range item.RegisterToPatient {
	//	item.RegisterToPatientForDelete = append(item.RegisterToPatientForDelete, item.RegisterToPatient[i].ID)
	//}
}
