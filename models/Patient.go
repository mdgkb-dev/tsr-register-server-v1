package models

import (
	"encoding/json"
	"errors"
	"fmt"

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
	Disabilities                     Disabilities            `bun:"rel:has-many" json:"disabilities"`

	PatientDiagnosis          []*PatientDiagnosis `bun:"rel:has-many" json:"patientDiagnosis"`
	PatientDiagnosisForDelete []uuid.UUID         `bun:"-" json:"patientDiagnosisForDelete"`

	PatientDrugRegimen          []*PatientDrugRegimen `bun:"rel:has-many" json:"patientDrugRegimen"`
	PatientDrugRegimenForDelete []uuid.UUID           `bun:"-" json:"patientDrugRegimenForDelete"`

	PatientsRegisters PatientsRegisters `bun:"rel:has-many" json:"patientsRegisters"`
	Commissions       Commissions       `bun:"rel:has-many" json:"commissions"`

	PatientsResearchesPools          PatientsResearchesPools `bun:"rel:has-many" json:"patientsResearchesPools"`
	PatientsResearchesPoolsForDelete []uuid.UUID             `bun:"-" json:"patientsResearchesPoolsForDelete"`

	PatientsResearches          PatientsResearches `bun:"rel:has-many" json:"patientsResearches"`
	PatientsResearchesForDelete []uuid.UUID        `bun:"-" json:"patientsResearchesForDelete"`
	PatientHistories            PatientHistories   `bun:"rel:has-many" json:"patientHistories"`

	FullName  string `bun:"-" json:"fullName"`
	IsMale    string `bun:"-" json:"isMale"`
	DateBirth string `bun:"-" json:"dateBirth"`

	Anamneses Anamneses `bun:"rel:has-many" json:"anamneses"`
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
}

func (item *Patient) SetDeleteIDForChildren() {
	for i := range item.PatientsRepresentatives {
		item.PatientsRepresentativesForDelete = append(item.PatientsRepresentativesForDelete, item.PatientsRepresentatives[i].ID)
	}
	for i := range item.PatientDrugRegimen {
		item.PatientDrugRegimenForDelete = append(item.PatientDrugRegimenForDelete, item.PatientDrugRegimen[i].ID)
	}
}

func (item *Patient) GetMaxResearchesResultsCount() int {
	maxCount := 0
	for _, patientResearch := range item.PatientsResearches {
		resultsCount := len(patientResearch.ResearchResults)
		if resultsCount > maxCount {
			maxCount = resultsCount
		}
	}

	return maxCount
}

func (items Patients) GetExportData(researches Researches) ([][]interface{}, error) {
	fmt.Println("НАЧАЛО")
	dataLines := make([][]interface{}, 0)
	for _, patient := range items {
		patientData, err := patient.GetExportData(researches)
		if err != nil {
			return nil, err
		}
		fmt.Println("FIRST PATIENT LEN", len(patientData))
		dataLines = append(dataLines, patientData...)
	}
	return dataLines, nil
}

func (item *Patient) GetExportData(researches Researches) ([][]interface{}, error) {
	patientData := make([][]interface{}, 0)
	patientData = append(patientData, []interface{}{})
	patientData[0] = append(patientData[0], item.Human.GetFullName())
	if item.Human.DateBirth != nil {
		patientData[0] = append(patientData[0], item.Human.DateBirth.Format("02.01.2006"))
	} else {
		patientData[0] = append(patientData[0], "")
	}
	// colIndex := 2
	//patientData[0][0] = item.Human.GetFullName()
	//patientData[0][1] = item.Human.DateBirth.Format("02.01.2006")
	startColIdx := 0
	for researchIdx, research := range researches {
		// fmt.Println("startColIdx", startColIdx)
		// if researchIndex == 0 && item.Human != nil {

		// }
		patientResearch, err := item.GetPatientResearch(research.ID)
		if err != nil {
			return nil, err
		}
		patientResearchResults, err := patientResearch.GetExportData(research)
		if err != nil {
			return nil, err
		}
		// fmt.Println(len(patientResearchResults))
		// fmt.Println(patientResearchResults)
		// fmt.Println(patientData)

		for i, result := range patientResearchResults {
			// Добавили строку
			if i >= len(patientData) {
				patientData = append(patientData, make([]interface{}, 0))
			}
			if i > 0 && researchIdx == 0 {
				fmt.Println("Начало заполнения", researchIdx, i, patientData[i])
				patientData[i] = append(patientData[i], []interface{}{""})
				patientData[i] = append(patientData[i], []interface{}{""})
			}
			fmt.Println("Заполнили пустые", researchIdx, i, patientData[i])
			if startColIdx > 0 && len(patientData[i]) < startColIdx {
				// fmt.Println("added", startColIdx)
				patientData[i] = append(patientData[i], []interface{}{""})
				patientData[i] = append(patientData[i], []interface{}{""})
				patientData[i] = append(patientData[i], make([]interface{}, startColIdx)...)
			}
			for _, answer := range result {
				patientData[i] = append(patientData[i], answer)
			}
			fmt.Println("Заполнили ответы", researchIdx, i, patientData[i])
		}
		startColIdx += len(patientResearchResults[0])
		// fmt.Println("Ответов", len(patientResearchResults[0]))

	}
	//fmt.Println("patient")
	return patientData, nil
}

func (item *Patient) GetPatientResearch(researchID uuid.NullUUID) (res *PatientResearch, err error) {
	for _, patientResearch := range item.PatientsResearches {
		if researchID.UUID.String() == patientResearch.ResearchID.UUID.String() {
			res = patientResearch
			break
		}
	}
	if res == nil {
		//return nil, errors.New("у пациента отсутствует исследование")
		return &PatientResearch{}, nil
	}
	return res, nil
}

type PatientsExport struct {
	IDPool          []string `json:"ids"`
	WithAge         bool     `json:"withAge"`
	CountAverageAge bool     `json:"countAverageAge"`
}

const patientsExportOptionsKey = "patient"

func (item *PatientsExport) ParseExportOptions(options map[string]map[string]interface{}) error {
	patientsOptions, ok := options[patientsExportOptionsKey]
	if !ok {
		return errors.New("not find patients")
	}
	jsonbody, err := json.Marshal(patientsOptions[patientsExportOptionsKey])
	if err != nil {
		return errors.New("parse error")
	}

	if err := json.Unmarshal(jsonbody, &item); err != nil {
		return errors.New("parse error")
	}

	return nil
}
