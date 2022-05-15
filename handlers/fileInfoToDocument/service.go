package fileInfoForDocument

import (
	"mdgkb/tsr-tegister-server-v1/handlers/fileInfo"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items []*models.FileInfoToDocument) error {
	if len(items) == 0 {
		return nil
	}
	err := fileInfo.CreateService(s.repository.getDB(), s.helper).CreateMany(models.GetFileInfoFileInfoToDocument(items))
	if err != nil {
		return err
	}
	for i := range items {
		items[i].FileInfoID = items[i].FileInfo.ID
	}
	err = s.repository.createMany(items)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items []*models.FileInfoToDocument) error {
	if len(items) == 0 {
		return nil
	}
	err := fileInfo.CreateService(s.repository.getDB(), s.helper).UpsertMany(models.GetFileInfoFileInfoToDocument(items))
	if err != nil {
		return err
	}
	for i := range items {
		items[i].FileInfoID = items[i].FileInfo.ID
	}
	return s.repository.upsertMany(items)
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
