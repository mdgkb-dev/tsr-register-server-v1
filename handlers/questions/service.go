package questions

import (
	"mdgkb/tsr-tegister-server-v1/handlers/answervariants"
	"mdgkb/tsr-tegister-server-v1/handlers/questionexamples"
	"mdgkb/tsr-tegister-server-v1/handlers/questionmeasures"
	"mdgkb/tsr-tegister-server-v1/handlers/registerpropertyvariants"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Service) Create(item *models.Question) error {
	return s.repository.Create(item)
}

func (s *Service) GetAll() (models.QuestionsWithCount, error) {
	items, err := s.repository.GetAll()
	if err != nil {
		return items, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.Question, error) {
	item, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.Question) error {
	return s.repository.Update(item)
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *Service) UpsertMany(items models.Questions) error {
	if len(items) == 0 {
		return nil
	}

	err := s.repository.upsertMany(items)
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
	registerPropertyVariantsService := registerpropertyvariants.CreateService(s.helper)
	err = registerPropertyVariantsService.UpsertMany(items.GetRegisterPropertyVariants())
	if err != nil {
		return err
	}
	err = registerPropertyVariantsService.DeleteMany(items.GetRegisterPropertyVariantsForDelete())
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

func (s *Service) SetQueryFilter(c *gin.Context) (err error) {
	err = s.repository.SetQueryFilter(c)
	return err
}
