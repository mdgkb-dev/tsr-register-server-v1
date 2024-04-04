package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SelectedAnswerVariant struct {
	bun.BaseModel `bun:"selected_answer_variants,alias:selected_answer_variants"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Answer        *Answer       `bun:"rel:belongs-to" json:"answer"`
	AnswerID      uuid.NullUUID `bun:"type:uuid" json:"answerId"`

	AnswerVariant   *AnswerVariant `bun:"rel:belongs-to" json:"answerVariant"`
	AnswerVariantID uuid.NullUUID  `bun:"type:uuid" json:"answerVariantId"`
	PatientID       uuid.NullUUID  `bun:"type:uuid" json:"patientId"`
}

type SelectedAnswerVariants []*SelectedAnswerVariant

func (item *SelectedAnswerVariant) SetIDForChildren() {
	//if len(item.RegisterPropertyOthers) == 0 {
	//	return
	//}
	//for i := range item.RegisterPropertyOthers {
	//	item.RegisterPropertyOthers[i].AnswerVariantID = item.ID
	//}
}

func (items SelectedAnswerVariants) SetIDForChildren() {
	if len(items) == 0 {
		return
	}
	for i := range items {
		items[i].SetIDForChildren()
	}
}
