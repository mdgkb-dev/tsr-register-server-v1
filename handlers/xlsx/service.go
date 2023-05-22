package xlsx

import (
	"mdgkb/tsr-tegister-server-v1/handlers/registerquery"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) GetRegisterQuery(id string) (*models.ResearchQuery, error) {
	return registerquery.CreateService(s.helper).Get(id)
}
