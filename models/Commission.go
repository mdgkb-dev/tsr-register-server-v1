package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Commission struct {
	bun.BaseModel      `bun:"commissions,alias:commissions"`
	ID                 uuid.NullUUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Date               *time.Time        `bun:"item_date" json:"date"`
	StartDate          *time.Time        `json:"startDate"`
	EndDate            *time.Time        `json:"endDate"`
	Volume             string            `json:"volume"`
	Number             int               `bun:",autoincrement,notnull," json:"number"`
	Patient            *Patient          `bun:"rel:belongs-to" json:"patient"`
	PatientID          uuid.NullUUID     `bun:"type:uuid" json:"patientId"`
	CommissionStatus   *CommissionStatus `bun:"rel:belongs-to" json:"commissionStatus"`
	CommissionStatusID uuid.NullUUID     `bun:"type:uuid" json:"commissionStatusId"`
	PatientDiagnosis   *PatientDiagnosis `bun:"rel:belongs-to" json:"patientDiagnosis"`
	PatientDiagnosisID uuid.NullUUID     `bun:"type:uuid" json:"patientDiagnosisId"`
	DrugRegimen        *DrugRegimen      `bun:"rel:belongs-to" json:"drugRegimen"`
	DrugRegimenID      uuid.NullUUID     `bun:"type:uuid" json:"drugRegimenId"`

	Drug   *Drug         `bun:"rel:belongs-to" json:"drug"`
	DrugID uuid.NullUUID `bun:"type:uuid" json:"drugId"`

	DzmAnswerFile               *FileInfo                   `bun:"rel:belongs-to" json:"dzmAnswerFile"`
	DzmAnswerFileID             uuid.NullUUID               `bun:"type:uuid" json:"dzmAnswerFileId"`
	DzmAnswerComment            string                      `json:"dzmAnswerComment"`
	CommissionsDoctors          CommissionsDoctors          `bun:"rel:has-many" json:"commissionsDoctors"`
	CommissionsDrugApplications CommissionsDrugApplications `bun:"rel:has-many" json:"commissionsDrugApplications"`
}

type Commissions []*Commission
type CommissionsWithCount struct {
	Commissions Commissions `json:"items"`
	Count       int         `json:"count"`
}

func (item *Commission) SetIDForChildren() {
	for i := range item.CommissionsDoctors {
		item.CommissionsDoctors[i].CommissionID = item.ID
	}
}
