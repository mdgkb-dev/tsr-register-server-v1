package drug

import (
	"mdgkb/tsr-tegister-server-v1/handlers/drugRegimen"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.Drug) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = drugRegimen.CreateService(s.repository.getDB()).CreateMany(item.DrugRegimens)
	if err != nil {
		return err
	}
	return nil

}

func (s *Service) GetAll() ([]*models.Drug, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.Drug, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.Drug) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	drugRegimenService := drugRegimen.CreateService(s.repository.getDB())
	err = drugRegimenService.UpsertMany(item.DrugRegimens)
	if err != nil {
		return err
	}
	err = drugRegimenService.DeleteMany(item.DrugRegimensForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}
