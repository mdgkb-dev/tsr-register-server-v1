package registerQuery

import (
	"mdgkb/tsr-tegister-server-v1/handlers/registerQueryToRegisterProperty"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(query *models.RegisterQuery) error {
	err := s.repository.create(query)

	if err != nil {
		return err
	}

	query.SetIdForChildren()
	err = registerQueryToRegisterProperty.CreateService(s.repository.getDB()).CreateMany(query.RegisterQueryToRegisterProperty)
	return err
}

func (s *Service) GetAll() (models.RegisterQueries, error) {
	queries, err := s.repository.getAll()

	if err != nil {
		return nil, err
	}

	return queries, nil
}

func (s *Service) Get(id *string) (*models.RegisterQuery, error) {
	query, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return query, nil
}

func (s *Service) Update(query *models.RegisterQuery) error {
	err := s.repository.update(query)

	if err != nil {
		return err
	}

	query.SetIdForChildren()
	registerQueryToRegisterPropertyService := registerQueryToRegisterProperty.CreateService(s.repository.getDB())
	err = registerQueryToRegisterPropertyService.UpsertMany(query.RegisterQueryToRegisterProperty)

	if err != nil {
		return err
	}

	err = registerQueryToRegisterPropertyService.DeleteMany(query.RegisterQueryToRegisterPropertyForDelete)
	return err
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) Execute(id string) ([]map[string]interface{}, error) {
	registerQuery, err := s.repository.get(&id)
	if err != nil {
		return nil, err
	}

	res, err := s.repository.execute(registerQuery)
	if err != nil {
		return nil, err
	}
	return res, nil
}
