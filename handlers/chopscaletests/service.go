package chopscaletests

import (
	"mdgkb/tsr-tegister-server-v1/handlers/chopscaletestresults"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items models.ChopScaleTests) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	err = chopscaletestresults.CreateService(s.helper).CreateMany(items.GetChopScaleTestResults())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.ChopScaleTests) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	chopScaleTestResultsService := chopscaletestresults.CreateService(s.helper)
	err = chopScaleTestResultsService.UpsertMany(items.GetChopScaleTestResults())
	if err != nil {
		return err
	}
	err = chopScaleTestResultsService.DeleteMany(items.GetChopScaleTestResultsForDelete())
	return err
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
