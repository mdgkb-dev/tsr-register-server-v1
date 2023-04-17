package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PatientAnswerComment struct {
	bun.BaseModel            `bun:"patient_answer_comments,alias:patient_answer_comments"`
	ID                       uuid.UUID       `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Value                    string          `json:"value"`
	AnswerComment            *AnswerComment  `bun:"rel:belongs-to" json:"answerComment"`
	AnswerCommentID          uuid.UUID       `bun:"type:uuid" json:"answerCommentId"`
	PatientResearchSection   *ResearchResult `bun:"rel:belongs-to" json:"patientResearchSection"`
	PatientResearchSectionID uuid.UUID       `bun:"type:uuid" json:"patientResearchSectionId"`
}

type PatientAnswerComments []*PatientAnswerComment
