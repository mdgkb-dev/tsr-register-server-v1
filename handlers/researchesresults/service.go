package researchesresults

import (
	"mdgkb/tsr-tegister-server-v1/handlers/answers"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.ResearchResult) error {
	err := s.repository.Create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	answersService := answers.CreateService(s.helper)
	err = answersService.UpsertMany(item.Answers)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.ResearchResult) error {
	err := s.repository.Update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	answersService := answers.CreateService(s.helper)
	err = answersService.UpsertMany(item.Answers)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.ResearchResultsWithCount, error) {
	return s.repository.GetAll()
}

func (s *Service) Get(slug string) (*models.ResearchResult, error) {
	item, err := s.repository.Get(slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *Service) SetQueryFilter(c *gin.Context) (err error) {
	err = s.repository.SetQueryFilter(c)
	return err
}
