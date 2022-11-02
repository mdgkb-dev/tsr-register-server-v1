package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
)

type RegisterPropertyToPatient struct {
	bun.BaseModel `bun:"register_property_to_patient,alias:register_property_to_patient"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	ValueString string    `json:"valueString"`
	ValueNumber float32   `json:"valueNumber"`
	ValueDate   time.Time `json:"valueDate"`
	ValueOther  string    `json:"valueOther"`

	RegisterPropertiesToPatientsToFileInfos          RegisterPropertiesToPatientsToFileInfos `bun:"rel:has-many" json:"registerPropertiesToPatientsToFileInfos"`
	RegisterPropertiesToPatientsToFileInfosForDelete []uuid.UUID                             `bun:"-" json:"registerPropertiesToPatientsToFileInfosForDelete"`

	RegisterPropertyRadio   *RegisterPropertyRadio `bun:"rel:belongs-to" json:"registerPropertyRadio"`
	RegisterPropertyRadioID uuid.NullUUID          `bun:"type:uuid" json:"registerPropertyRadioId"`

	RegisterPropertyMeasure   *RegisterPropertyMeasure `bun:"rel:belongs-to" json:"registerPropertyMeasure"`
	RegisterPropertyMeasureID uuid.NullUUID            `bun:"type:uuid" json:"registerPropertyMeasureId"`

	RegisterProperty   *RegisterProperty `bun:"rel:belongs-to" json:"registerProperty"`
	RegisterPropertyID uuid.UUID         `bun:"type:uuid" json:"registerPropertyId"`

	RegisterPropertyVariant   *RegisterPropertyVariant `bun:"rel:belongs-to" json:"registerPropertyVariant"`
	RegisterPropertyVariantID uuid.NullUUID            `bun:"type:uuid" json:"registerPropertyVariantId"`

	RegisterGroupToPatient   *RegisterGroupToPatient `bun:"rel:belongs-to" json:"registerGroupToPatient"`
	RegisterGroupToPatientID uuid.UUID               `bun:"type:uuid" json:"registerGroupToPatientID"`
}

type RegisterPropertiesToPatients []*RegisterPropertyToPatient

func (item *RegisterPropertyToPatient) SetFilePath(fileID *string) *string {
	for i := range item.RegisterPropertiesToPatientsToFileInfos {
		if item.RegisterPropertiesToPatientsToFileInfos[i].FileInfo.ID.String() == *fileID {
			item.RegisterPropertiesToPatientsToFileInfos[i].FileInfo.FileSystemPath = uploadHelper.BuildPath(fileID)
			return &item.RegisterPropertiesToPatientsToFileInfos[i].FileInfo.FileSystemPath
		}
	}
	return nil
}

func (item *RegisterPropertyToPatient) SetIDForChildren() {
	if len(item.RegisterPropertiesToPatientsToFileInfos) > 0 {
		for i := range item.RegisterPropertiesToPatientsToFileInfos {
			item.RegisterPropertiesToPatientsToFileInfos[i].RegisterPropertyToPatientID = item.ID
		}
	}
}

func (items RegisterPropertiesToPatients) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}
}

func (items RegisterPropertiesToPatients) GetRegisterPropertiesToPatientsToFileInfos() RegisterPropertiesToPatientsToFileInfos {
	itemsForGet := make(RegisterPropertiesToPatientsToFileInfos, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertiesToPatientsToFileInfos...)
	}

	return itemsForGet
}

func (items RegisterPropertiesToPatients) GetRegisterPropertiesToPatientsToFileInfosForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertiesToPatientsToFileInfosForDelete...)
	}

	return itemsForGet
}
