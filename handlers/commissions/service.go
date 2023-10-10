package commissions

import (
	"mdgkb/tsr-tegister-server-v1/handlers/commissionsdoctors"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.Commission) error {
	err := s.repository.Create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = commissionsdoctors.CreateService(s.helper).UpsertMany(item.CommissionsDoctors)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) GetAll() (models.CommissionsWithCount, error) {
	return s.repository.GetAll()
}

func (s *Service) Get(id string) (*models.Commission, error) {
	item, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.Commission) error {
	err := s.repository.Update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = commissionsdoctors.CreateService(s.helper).UpsertMany(item.CommissionsDoctors)
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
