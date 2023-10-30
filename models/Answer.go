package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Answer struct {
	bun.BaseModel `bun:"answers,alias:answers"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	ValueString     string         `json:"valueString"`
	ValueNumber     float32        `json:"valueNumber"`
	ValueDate       time.Time      `json:"valueDate"`
	ValueOther      string         `json:"valueOther"`
	AnswerVariant   *AnswerVariant `bun:"rel:belongs-to" json:"answerVariant"`
	AnswerVariantID uuid.NullUUID  `bun:"type:uuid" json:"answerVariantId"`

	ResearchResult   *ResearchResult `bun:"rel:belongs-to" json:"researchResult"`
	ResearchResultID uuid.NullUUID   `bun:"type:uuid" json:"researchResultId"`
	Filled           bool            `json:"filled"`

	Question   *Question     `bun:"rel:belongs-to" json:"question"`
	QuestionID uuid.NullUUID `bun:"type:uuid" json:"questionId"`
	PatientID  uuid.NullUUID `bun:"type:uuid" json:"patientId"`

	QuestionVariant   *QuestionVariant `bun:"rel:belongs-to" json:"questionVariant"`
	QuestionVariantID uuid.NullUUID    `bun:"type:uuid" json:"questionVariantId"`

	SelectedAnswerVariants          SelectedAnswerVariants `bun:"rel:has-many" json:"selectedAnswerVariants"`
	SelectedAnswerVariantsForDelete []uuid.UUID            `bun:"-" json:"selectedAnswerVariantsForDelete"`

	AnswerFiles          AnswerFiles `bun:"rel:has-many" json:"answerFiles"`
	AnswerFilesForDelete []uuid.UUID `bun:"-" json:"answerFilesForDelete"`
}

type Answers []*Answer

func (items Answers) SetFilePath(fileID *string) *string {
	for i := range items {
		path := items[i].SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	return nil
}

func (item *Answer) SetFilePath(fileID *string) *string {
	for i := range item.AnswerFiles {
		path := item.AnswerFiles[i].SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	return nil
}

func (items Answers) GetSelectedAnswerVariants() SelectedAnswerVariants {
	itemsForGet := make(SelectedAnswerVariants, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].SelectedAnswerVariants...)
	}

	return itemsForGet
}

func (items Answers) GetSelectedAnswerVariantsForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].SelectedAnswerVariantsForDelete...)
	}

	return itemsForGet
}
func (item *Answer) SetIDForChildren() {
	for i := range item.AnswerFiles {
		item.AnswerFiles[i].AnswerID = item.ID
	}
	for i := range item.SelectedAnswerVariants {
		item.AnswerFiles[i].AnswerID = item.ID
	}
}

func (items Answers) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}

	//if len(item.RegisterGroups) > 0 {
	//	for i := range item.RegisterGroups {
	//		item.RegisterGroups[i].ResearchPoolID = item.ID
	//	}
	//}
	//if len(item.RegisterDiagnosis) > 0 {
	//	for i := range item.RegisterDiagnosis {
	//		item.RegisterDiagnosis[i].ResearchPoolID = item.ID
	//	}
	//}
}

func (items Answers) GetAnswerFiles() AnswerFiles {
	itemsForGet := make(AnswerFiles, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].AnswerFiles...)
	}

	return itemsForGet
}

const (
	Yes    string = "Да"
	No     string = "Нет"
	NoData string = "Нет данных"
)

func (items Answers) GetAnswerFilesForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].AnswerFilesForDelete...)
	}

	return itemsForGet
}

func (item *Answer) GetData(q *Question) interface{} {
	if q.ValueType.IsString() || q.ValueType.IsText() {
		return item.ValueString
	}
	if q.ValueType.IsNumber() {
		return item.ValueNumber
	}
	if q.ValueType.IsDate() {
		return item.ValueDate
	}
	if q.ValueType.IsRadio() {
		res := No
		for _, radio := range q.AnswerVariants {
			if radio.ID == item.AnswerVariantID {
				res = radio.Name
				break
			}
		}
		return res
	}
	if q.ValueType.IsSet() {
		res := ""
		for _, v := range item.SelectedAnswerVariants {
			res += v.AnswerVariant.Name + "; "
		}
		// for _, radio := range q.AnswerVariants {
		// if radio.ID == item.AnswerVariantID {
		// 	res = radio.Name
		// 	break
		// }
		// }
		return res
	}
	return ""
}

func (item *Answer) GetAggregateExistingData() bool {
	if item.Question.ValueType.IsString() || item.Question.ValueType.IsText() {
		return len(item.ValueString) > 0
	}
	if item.Question.ValueType.IsNumber() {
		return item.ValueNumber > 0
	}
	if item.Question.ValueType.IsDate() {
		return !item.ValueDate.IsZero()
	}
	if item.Question.ValueType.IsRadio() {
		if item.AnswerVariantID.Valid {
			return true
		}
	}
	return false
}

func (item *Answer) AnswerVariantSelected(variantID uuid.NullUUID) string {
	res := No
	for _, selectedVariant := range item.SelectedAnswerVariants {
		if selectedVariant.AnswerVariantID == variantID {
			res = Yes
			break
		}
	}
	return res
}
