package xlsx

import (
	"mdgkb/tsr-tegister-server-v1/handlers/registerQuery"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) GetRegisterQuery(id *string) (*models.RegisterQuery, error) {
	return registerQuery.CreateService(s.repository.getDB()).Get(id)
}
