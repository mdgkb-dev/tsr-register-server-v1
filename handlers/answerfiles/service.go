package answerfiles

import (
	"mdgkb/tsr-tegister-server-v1/handlers/fileinfos"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) Create(item *models.AnswerFile) error {
	return s.repository.create(item)
}

func (s *Service) GetAll() ([]*models.AnswerFile, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.AnswerFile, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.AnswerFile) error {
	return s.repository.update(item)
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) UpsertMany(items models.AnswerFiles) error {
	if len(items) == 0 {
		return nil
	}

	fileInfosService := fileinfos.CreateService(s.helper)
	err := fileInfosService.UpsertMany(items.GetFileInfos())
	if err != nil {
		return err
	}

	items.SetForeignKeys()

	err = s.repository.upsertMany(items)
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
