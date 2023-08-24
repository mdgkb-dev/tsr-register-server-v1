package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MkbQuestionDomain struct {
	bun.BaseModel `bun:"mkb_questions_domains,alias:mkb_questions_domains"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	QuestionID    uuid.NullUUID `bun:"type:uuid" json:"questionId"`
	Question      *Question     `bun:"-" json:"question"`

	DomainID uuid.NullUUID `bun:"type:uuid" json:"domainId"`
	Domain   *Domain       `bun:"-" json:"domain"`
}

type MkbQuestionsDomains []*MkbQuestionDomain

type MkbQuestionsDomainsWithCount struct {
	MkbQuestionsDomains MkbQuestionsDomains `json:"items"`
	Count               int                 `json:"count"`
}
