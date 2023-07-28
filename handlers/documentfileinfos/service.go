package documentfileinfos

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.DocumentFileInfo) error {
	err := s.repository.Create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.DocumentFileInfo) error {
	err := s.repository.Update(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.DocumentFileInfosWithCount, error) {
	return s.repository.GetAll()
}

func (s *Service) Get(slug string) (*models.DocumentFileInfo, error) {
	item, err := s.repository.Get(slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *Service) UpsertMany(items models.DocumentFileInfos) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.UpsertMany(items)
}

func (s *Service) SetQueryFilter(c *gin.Context) (err error) {
	err = s.repository.SetQueryFilter(c)
	return err
}
