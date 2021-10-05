package human

import (
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *HistoryService) Create(item *models.HumanHistory) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}
