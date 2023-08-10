package models

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Answer struct {
	bun.BaseModel `bun:"answers,alias:answers"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	//Order         int       `bun:"item_order" json:"order"`

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
	//RegisterPropertiesToPatientsToFileInfos          RegisterPropertiesToPatientsToFileInfos `bun:"rel:has-many" json:"registerPropertiesToPatientsToFileInfos"`
	//RegisterPropertiesToPatientsToFileInfosForDelete []uuid.UUID                             `bun:"-" json:"registerPropertiesToPatientsToFileInfosForDelete"`

	//RegisterPropertyMeasure   *RegisterPropertyMeasure `bun:"rel:belongs-to" json:"registerPropertyMeasure"`
	//RegisterPropertyMeasureID uuid.NullUUID            `bun:"type:uuid" json:"registerPropertyMeasureId"`

	//RegisterPropertyVariant   *RegisterPropertyVariant `bun:"rel:belongs-to" json:"registerPropertyVariant"`
	//RegisterPropertyVariantID uuid.NullUUID            `bun:"type:uuid" json:"registerPropertyVariantId"`

	//RegisterGroupToPatient   *RegisterGroupToPatient `bun:"rel:belongs-to" json:"registerGroupToPatient"`
	//RegisterGroupToPatientID uuid.UUID               `bun:"type:uuid" json:"registerGroupToPatientID"`
}

type Answers []*Answer

//
//func (item *RegisterPropertyToPatient) SetFilePath(fileID *string) *string {
//	for i := range item.RegisterPropertiesToPatientsToFileInfos {
//		if item.RegisterPropertiesToPatientsToFileInfos[i].FileInfo.ID.String() == *fileID {
//			item.RegisterPropertiesToPatientsToFileInfos[i].FileInfo.FileSystemPath = uploadHelper.BuildPath(fileID)
//			return &item.RegisterPropertiesToPatientsToFileInfos[i].FileInfo.FileSystemPath
//		}
//	}
//	return nil
//}
//
//func (item *RegisterPropertyToPatient) SetIDForChildren() {
//	if len(item.RegisterPropertiesToPatientsToFileInfos) > 0 {
//		for i := range item.RegisterPropertiesToPatientsToFileInfos {
//			item.RegisterPropertiesToPatientsToFileInfos[i].RegisterPropertyToPatientID = item.ID
//		}
//	}
//}
//
//func (items Answers) SetIDForChildren() {
//	for i := range items {
//		items[i].SetIDForChildren()
//	}
//}
//
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

//
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

func (item *Answer) GetData(prop *Question) string {
	if prop.ValueType.IsString() || prop.ValueType.IsText() {
		return item.ValueString
	}
	if prop.ValueType.IsNumber() {
		return strconv.Itoa(int(item.ValueNumber))
	}
	if prop.ValueType.IsDate() {
		return item.ValueDate.String()
	}
	if prop.ValueType.IsRadio() {
		res := No
		for _, radio := range prop.AnswerVariants {
			//if prop.Name == "Атопический дерматит" {
			//	fmt.Println(radio.ID.UUID, radio.Name, item.ID.UUID, item.QuestionID.UUID, prop.ID.UUID)
			//}
			if radio.ID == item.AnswerVariantID {
				res = radio.Name
				break
			}
		}
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
