package documents

import (
	"mdgkb/tsr-tegister-server-v1/handlers/documentfieldvalues"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.Document) error {
	err := s.repository.Create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = documentfieldvalues.CreateService(s.helper).UpsertMany(item.DocumentFieldValues)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Update(item *models.Document) error {
	err := s.repository.Update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	return nil
}

func (s *Service) GetAll() (models.DocumentsWithCount, error) {
	return s.repository.GetAll()
}

func (s *Service) Get(slug string) (*models.Document, error) {
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
