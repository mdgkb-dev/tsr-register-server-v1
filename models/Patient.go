package models

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Pramod-Devireddy/go-exprtk"
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

func (item *Patient) GetXlsxData(research *Research) (results [][]interface{}, err error) {
	m := exprtk.NewExprtk()
	defer m.Delete()
	results = make([][]interface{}, 0)
	patientResearch, err := item.GetPatientResearch(research.ID.UUID.String())
	if err != nil {
		return nil, err
	}

	for resultN, researchResult := range patientResearch.ResearchResults {
		variables := make(map[string]interface{})
		results = append(results, []interface{}{})
		results[resultN] = append(results[resultN], researchResult.Date.Format("02.01.2006"))

		if research.WithScores {
			sum := 0
			for _, q := range research.Questions {
				sum += researchResult.GetScores(q)
			}
			results[resultN] = append(results[resultN], strconv.Itoa(sum))
			continue
		}

		for _, q := range research.Questions {
			answer := researchResult.GetData(q)
			results[resultN] = append(results[resultN], answer)
			variables[q.Code] = answer
			//i, e := strconv.Atoi(answer)
			//if e == nil {
			//	variables[q.Code] = i
			//}
		}

		for _, f := range research.Formulas {
			if !f.Xlsx {
				continue
			}
			m.SetExpression(f.Formula)
			for k := range variables {
				m.AddDoubleVariable(k)
			}
			err = m.CompileExpression()
			if err != nil {
				fmt.Println(err)
			}
			//for k, v := range variables {
			//i, err := strconv.Atoi(v.(string))
			//if err != nil {
			//	fmt.Println(err)
			//}
			//m.SetDoubleVariableValue(k, float64(i))
			//}
			//
			value := m.GetEvaluatedValue()
			//answer := researchResult.GetData(q)
			results[resultN] = append(results[resultN], fmt.Sprintf("%.2f", value))
			fmt.Println(results[resultN])
		}
	}
	return results, nil
}

func (item *Patient) GetPatientResearch(researchID string) (res *PatientResearch, err error) {
	for _, patientResearch := range item.PatientsResearches {
		if researchID == patientResearch.ResearchID.UUID.String() {
			res = patientResearch
			break
		}
	}
	if res == nil {
		return nil, errors.New("у пациента отсутствует исследование")
	}
	return res, nil
}
