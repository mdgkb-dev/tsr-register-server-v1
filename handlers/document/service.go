package document

import (
	documentFieldValues "mdgkb/tsr-tegister-server-v1/handlers/documentFieldValue"
	fileInfoForDocument "mdgkb/tsr-tegister-server-v1/handlers/fileInfoToDocument"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items []*models.Document) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	err = documentFieldValues.CreateService(s.helper).CreateMany(models.GetDocumentsFiledValues(items))
	if err != nil {
		return err
	}

	err = fileInfoForDocument.CreateService(s.helper).CreateMany(models.GetFileInfoToDocument(items))
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpsertMany(items []*models.Document) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	err = documentFieldValues.CreateService(s.helper).UpsertMany(models.GetDocumentsFiledValues(items))
	if err != nil {
		return err
	}
	fileInfoForDocumentService := fileInfoForDocument.CreateService(s.helper)
	err = fileInfoForDocumentService.UpsertMany(models.GetFileInfoToDocument(items))
	if err != nil {
		return err
	}
	err = fileInfoForDocumentService.DeleteMany(models.GetFileInfoToDocumentForDelete(items))
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
