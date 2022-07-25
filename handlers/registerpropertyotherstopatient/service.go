package registerpropertyotherstopatient

import (
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) UpsertMany(items models.RegisterPropertyOthersToPatient) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.upsertMany(items)
}
