package questions

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/handlers/answervariants"
	"mdgkb/tsr-tegister-server-v1/handlers/questionexamples"
	"mdgkb/tsr-tegister-server-v1/handlers/questionmeasures"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) Create(c context.Context, item *models.Question) error {
	return s.repository.Create(c, item)
}

func (s *Service) GetAll(c context.Context) (models.QuestionsWithCount, error) {
	items, err := s.repository.GetAll(c)
	if err != nil {
		return items, err
	}
	return items, nil
}

func (s *Service) Get(c context.Context, id string) (*models.Question, error) {
	item, err := s.repository.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(c context.Context, item *models.Question) error {
	return s.repository.Update(c, item)
}

func (s *Service) Delete(c context.Context, id string) error {
	return s.repository.Delete(c, id)
}

func (s *Service) UpsertMany(c context.Context, items models.Questions) error {
	if len(items) == 0 {
		return nil
	}

	err := s.repository.upsertMany(c, items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	registerPropertyRadioService := answervariants.CreateService(s.helper)
	err = registerPropertyRadioService.UpsertMany(items.GetRegisterPropertyRadios())
	if err != nil {
		return err
	}
	err = registerPropertyRadioService.DeleteMany(items.GetRegisterPropertyRadioForDelete())
	if err != nil {
		return err
	}

	registerPropertyExamplesService := questionexamples.CreateService(s.helper)
	err = registerPropertyExamplesService.UpsertMany(items.GetRegisterPropertyExamples())
	if err != nil {
		return err
	}
	err = registerPropertyExamplesService.DeleteMany(items.GetRegisterPropertyExamplesForDelete())
	if err != nil {
		return err
	}

	registerPropertyMeasuresService := questionmeasures.CreateService(s.helper)
	err = registerPropertyMeasuresService.UpsertMany(items.GetRegisterPropertyMeasures())
	if err != nil {
		return err
	}
	err = registerPropertyMeasuresService.DeleteMany(items.GetRegisterPropertyMeasuresForDelete())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteMany(c context.Context, idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(c, idPool)
}
