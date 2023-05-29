package researchquery

import (
	"mdgkb/tsr-tegister-server-v1/handlers/registerquerytoregisterproperty"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(query *models.ResearchQuery) error {
	err := s.repository.Create(query)

	if err != nil {
		return err
	}

	query.SetIDForChildren()
	err = registerquerytoregisterproperty.CreateService(s.helper).CreateMany(query.ResearchQueriesQuestions)
	return err
}

func (s *Service) GetAll() (models.ResearchQueriesWithCount, error) {
	return s.repository.GetAll()
}

func (s *Service) Get(id string) (*models.ResearchQuery, error) {
	return s.repository.Get(id)
}

func (s *Service) Update(query *models.ResearchQuery) error {
	err := s.repository.Update(query)

	if err != nil {
		return err
	}

	query.SetIDForChildren()
	registerQueryToRegisterPropertyService := registerquerytoregisterproperty.CreateService(s.helper)
	err = registerQueryToRegisterPropertyService.UpsertMany(query.ResearchQueriesQuestions)

	if err != nil {
		return err
	}

	//err = registerQueryToRegisterPropertyService.DeleteMany(query.RegisterQueryToRegisterPropertyForDelete)
	return err
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *Service) Execute(registerQuery *models.ResearchQuery) error {
	err := s.repository.Execute(registerQuery)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SetQueryFilter(c *gin.Context) (err error) {
	err = s.repository.SetQueryFilter(c)
	return err
}
