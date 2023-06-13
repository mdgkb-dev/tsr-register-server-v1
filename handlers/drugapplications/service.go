package drugapplications

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.DrugApplication) error {
	err := s.repository.Create(item)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) GetAll() (models.DrugApplicationsWithCount, error) {
	return s.repository.GetAll()
}

func (s *Service) Get(id string) (*models.DrugApplication, error) {
	item, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.DrugApplication) error {
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
