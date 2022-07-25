package documenttypes

import (
	documentTypeFields "mdgkb/tsr-tegister-server-v1/handlers/documenttypesfields"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.DocumentType) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = documentTypeFields.CreateService(s.helper).CreateMany(item.DocumentTypeFields)
	return err
}

func (s *Service) GetAll() ([]*models.DocumentType, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.DocumentType, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.DocumentType) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	documentTypeFieldsService := documentTypeFields.CreateService(s.helper)
	err = documentTypeFieldsService.UpsertMany(item.DocumentTypeFields)
	if err != nil {
		return err
	}
	if len(item.DocumentTypeFieldsForDelete) > 0 {
		err = documentTypeFieldsService.DeleteMany(item.DocumentTypeFieldsForDelete)
	}
	return err
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}
