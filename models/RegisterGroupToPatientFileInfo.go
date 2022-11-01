package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterPropertyToPatientToFileInfo struct {
	bun.BaseModel               `bun:"register_properties_to_patients_to_file_infos,alias:register_properties_to_patients_to_file_infos"`
	ID                          uuid.UUID                  `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	RegisterPropertyToPatient   *RegisterPropertyToPatient `bun:"rel:belongs-to" json:"registerPropertyToPatient"`
	RegisterPropertyToPatientID uuid.UUID                  `bun:"type:uuid" json:"registerPropertyToPatientId"`
	FileInfo                    *FileInfo                  `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoID                  uuid.UUID                  `bun:"type:uuid" json:"fileInfoId"`
}

type RegisterPropertiesToPatientsToFileInfos []*RegisterPropertyToPatientToFileInfo

func (item *RegisterPropertyToPatientToFileInfo) SetForeignKeys() {
	item.FileInfoID = item.FileInfo.ID
}

func (items RegisterPropertiesToPatientsToFileInfos) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}

func (items RegisterPropertiesToPatientsToFileInfos) GetFileInfos() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.FileInfo)
	}
	return itemsForGet
}
