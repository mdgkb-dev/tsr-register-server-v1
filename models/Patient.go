package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Patient struct {
	bun.BaseModel `bun:"patients,select:patients_view,alias:patients_view"`
	ModelInfo
	ID      uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Human   *Human        `bun:"rel:belongs-to" json:"human"`
	HumanID uuid.NullUUID `bun:"type:uuid" json:"humanId"`

	Region   *Region   `bun:"rel:belongs-to" json:"region"`
	RegionID uuid.UUID `bun:"type:uuid" json:"regionId"`

	PatientsRepresentatives          PatientsRepresentatives `bun:"rel:has-many" json:"patientsRepresentatives"`
	PatientsRepresentativesForDelete []uuid.UUID             `bun:"-" json:"patientsRepresentativesForDelete"`
	Disabilities                     Disabilities            `bun:"rel:has-many" json:"disabilities"`

	PatientDiagnosis          PatientDiagnoses `bun:"rel:has-many" json:"patientDiagnosis"`
	PatientDiagnosisForDelete []uuid.UUID      `bun:"-" json:"patientDiagnosisForDelete"`

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

	IsMoscow bool `json:"isMoscow"`
	Agreed   bool `json:"agreed"`

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
		item.PatientsRepresentativesForDelete = append(item.PatientsRepresentativesForDelete, item.PatientsRepresentatives[i].ID.UUID)
	}
	for i := range item.PatientDrugRegimen {
		item.PatientDrugRegimenForDelete = append(item.PatientDrugRegimenForDelete, item.PatientDrugRegimen[i].ID)
	}
}

func (items Patients) GetExportData(researches Researches) ([][]interface{}, Agregator, error) {
	dataLines := make([][]interface{}, 0)
	agregator := NewAgregator(researches.GetExportLen() + 2)
	for _, patient := range items {
		patientData, err := patient.GetExportData(researches)
		if err != nil {
			return nil, agregator, err
		}
		dataLines = append(dataLines, patientData...)

		for _, researchResult := range patientData {
			for answerIdx, answer := range researchResult {
				switch v := answer.(type) {
				case int:
					agregator.Sums[answerIdx] += float64(v)
					agregator.Count[answerIdx]++
				case float64:
					agregator.Sums[answerIdx] += v
					agregator.Count[answerIdx]++
				case float32:
					agregator.Sums[answerIdx] += float64(v)
					agregator.Count[answerIdx]++
				}
			}
		}
	}

	return dataLines, agregator, nil
}

func (item *Patient) GetExportData(researches Researches) ([][]interface{}, error) {
	patientData := make([][]interface{}, 0)
	patientData = append(patientData, []interface{}{})

	startColIdx := 0
	for researchIdx, research := range researches {
		patientResearch, err := item.GetPatientResearch(research.ID)
		if err != nil {
			return nil, err
		}

		patientResearchResults, err := patientResearch.GetExportData(research)
		if err != nil {
			return nil, err
		}

		if len(patientResearchResults) > len(patientData) {
			newLines := make([][]interface{}, len(patientResearchResults)-len(patientData))
			for i := range newLines {
				if len(patientData) > 0 {
					newLines[i] = append(newLines[i], make([]interface{}, len(patientData[0]))...)
				}
			}
			patientData = append(patientData, newLines...)
		}
		if researchIdx == 0 {
			startColIdx = 4
		} else {
			startColIdx = len(patientData[0])
		}

		for i := range patientData {
			if researchIdx == 0 {
				if i == 0 {
					patientData[i] = append(patientData[i], item.Human.GetFullName())
					if item.Human.DateBirth != nil {
						patientData[i] = append(patientData[i], item.Human.DateBirth.Format("02.01.2006"))
						patientData[i] = append(patientData[i], item.GetAge())
						patientData[i] = append(patientData[i], item.GetAgeInMonths())
					} else {
						patientData[i] = append(patientData[i], "")
						patientData[i] = append(patientData[i], "")
						patientData[i] = append(patientData[i], "")
					}
				} else {
					patientData[i] = append(patientData[i], []interface{}{""})
					patientData[i] = append(patientData[i], []interface{}{""})
					patientData[i] = append(patientData[i], []interface{}{""})
					patientData[i] = append(patientData[i], []interface{}{""})
				}
			}
			length := research.GetExportLen()
			patientData[i] = append(patientData[i], make([]interface{}, length)...)
		}

		for resultIDX, answers := range patientResearchResults {
			for answerIDX, answer := range answers {
				patientData[resultIDX][answerIDX+startColIdx] = answer
			}
		}
	}
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

func (item *Patient) GetAge() string {
	y, m := diff(*item.Human.DateBirth, time.Now())
	return fmt.Sprintf("%dг %d мес", y, m)
}

func (item *Patient) GetAgeInMonths() string {
	y, m := diff(*item.Human.DateBirth, time.Now())
	return fmt.Sprintf("%d", y*12+m)
}

func diff(a, b time.Time) (year, month int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, _ := a.Date()
	y2, M2, _ := b.Date()

	year = y2 - y1
	month = int(M2 - M1)
	if month < 0 {
		month += 12
		year--
	}

	return
}
