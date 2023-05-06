package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PatientHistory struct {
	bun.BaseModel `bun:"patient_histories,alias:patient_histories"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Patient       *Patient      `bun:"rel:belongs-to" json:"patients"`
	PatientID     uuid.NullUUID `bun:"type:uuid" json:"patientId"`
	User          *User         `bun:"rel:belongs-to" json:"user"`
	UserID        uuid.NullUUID `bun:"type:uuid" json:"userId"`
	CreatedAt     time.Time     `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`

	ActionType PatientHistoryActionType `json:"actionType"`
	ObjectCopy map[string]interface{}   `json:"objectCopy"`
}

type PatientHistories []*PatientHistory

type PatientHistoriesWithCount struct {
	PatientHistories PatientHistories `json:"items"`
	Count            int              `json:"count"`
}

type PatientHistoryActionType string

const (
	ActionCreate PatientHistoryActionType = "create"
	ActionUpdate PatientHistoryActionType = "update"
	ActionDelete PatientHistoryActionType = "delete"
)
