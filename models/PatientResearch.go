package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PatientResearch struct {
	//bun.BaseModel `bun:"patients_researches,select:patients_researches_view,alias:patients_researches_view"`
	bun.BaseModel `bun:"patients_researches,alias:patients_researches_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Research          *Research       `bun:"rel:belongs-to" json:"research"`
	ResearchID        uuid.NullUUID   `bun:"type:uuid" json:"researchId"`
	Patient           *Patient        `bun:"rel:belongs-to" json:"patients"`
	PatientID         uuid.NullUUID   `bun:"type:uuid" json:"patientId"`
	ResearchResults   ResearchResults `bun:"rel:has-many" json:"researchResults"`
	FillingPercentage uint            `json:"fillingPercentage"`
	Order             uint            `bun:"item_order" json:"order"`
}

type PatientsResearches []*PatientResearch

type PatientsResearchesWithCount struct {
	PatientsResearches PatientsResearches `json:"items"`
	Count              int                `json:"count"`
}

func (item *PatientResearch) GetExportData(research *Research) ([][]interface{}, error) {
	return item.ResearchResults.GetExportData(research)
}
