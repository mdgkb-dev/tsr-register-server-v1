package disability

import (
	"mdgkb/tsr-tegister-server-v1/handlers/edv"
	"mdgkb/tsr-tegister-server-v1/handlers/period"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items []*models.Disability) error {
	if len(items) == 0 {
		return nil
	}
	err := period.CreateService(s.repository.getDB()).CreateMany(models.GetPeriodsFromDisability(items))
	if err != nil {
		return err
	}
	models.SetPeriodIDToDisabilities(items)
	err = s.repository.createMany(items)
	if err != nil {
		return err
	}
	err = edv.CreateService(s.repository.getDB()).CreateMany(models.GetEdvs(items))
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items []*models.Disability) error {
	if len(items) == 0 {
		return nil
	}
	err := period.CreateService(s.repository.getDB()).UpsertMany(models.GetPeriodsFromDisability(items))
	if err != nil {
		return err
	}
	models.SetPeriodIDToDisabilities(items)
	err = s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	edvService := edv.CreateService(s.repository.getDB())
	err = edvService.UpsertMany(models.GetEdvs(items))
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
