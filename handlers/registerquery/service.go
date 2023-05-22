package registerquery

import (
	"mdgkb/tsr-tegister-server-v1/handlers/registerquerytoregisterproperty"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(query *models.ResearchQuery) error {
	err := s.repository.create(query)

	if err != nil {
		return err
	}

	query.SetIDForChildren()
	err = registerquerytoregisterproperty.CreateService(s.helper).CreateMany(query.ResearchQueriesQuestions)
	return err
}

func (s *Service) GetAll() (models.ResearchQueries, error) {
	queries, err := s.repository.getAll()

	if err != nil {
		return nil, err
	}

	return queries, nil
}

func (s *Service) Get(id string) (*models.ResearchQuery, error) {
	query, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return query, nil
}

func (s *Service) Update(query *models.ResearchQuery) error {
	err := s.repository.update(query)

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
	return s.repository.delete(id)
}

func (s *Service) Execute(registerQuery *models.ResearchQuery) error {
	err := s.repository.execute(registerQuery)
	if err != nil {
		return err
	}
	return nil
}
