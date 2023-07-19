package drugapplicationfiles

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.DrugApplicationFile) error {
	err := s.repository.Create(item)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) GetAll() (models.DrugApplicationFilesWithCount, error) {
	return s.repository.GetAll()
}

func (s *Service) Get(id string) (*models.DrugApplicationFile, error) {
	item, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.DrugApplicationFile) error {
	err := s.repository.Update(item)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *Service) SetQueryFilter(c *gin.Context) (err error) {
	err = s.repository.SetQueryFilter(c)
	return err
}
