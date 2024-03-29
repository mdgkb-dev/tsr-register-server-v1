package answers

import (
	"mdgkb/tsr-tegister-server-v1/handlers/answerfiles"
	"mdgkb/tsr-tegister-server-v1/handlers/selectedanswervariants"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) Create(item *models.Answer) error {
	return s.repository.create(item)
}

func (s *Service) GetAll() ([]*models.Answer, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.Answer, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.Answer) error {
	return s.repository.update(item)
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) UpsertMany(items models.Answers) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	selectedAnswerVariantsService := selectedanswervariants.CreateService(s.helper)
	err = selectedAnswerVariantsService.UpsertMany(items.GetSelectedAnswerVariants())
	if err != nil {
		return err
	}
	err = selectedAnswerVariantsService.DeleteMany(items.GetSelectedAnswerVariantsForDelete())
	if err != nil {
		return err
	}

	answerfilesService := answerfiles.CreateService(s.helper)
	err = answerfilesService.UpsertMany(items.GetAnswerFiles())
	if err != nil {
		return err
	}
	err = answerfilesService.DeleteMany(items.GetAnswerFilesForDelete())
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
