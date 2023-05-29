package xlsx

import (
	"mdgkb/tsr-tegister-server-v1/handlers/researchquery"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) GetRegisterQuery(id string) (*models.ResearchQuery, error) {
	return researchquery.CreateService(s.helper).Get(id)
}
