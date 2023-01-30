package hmfsescaletests

import (
	"mdgkb/tsr-tegister-server-v1/handlers/hmfsescaletestresults"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items models.HmfseScaleTests) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	err = hmfsescaletestresults.CreateService(s.helper).CreateMany(items.GetHmfseScaleTestResults())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.HmfseScaleTests) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	hmfseScaleTestResultsService := hmfsescaletestresults.CreateService(s.helper)
	err = hmfseScaleTestResultsService.UpsertMany(items.GetHmfseScaleTestResults())
	if err != nil {
		return err
	}
	err = hmfseScaleTestResultsService.DeleteMany(items.GetHmfseScaleTestResultsForDelete())
	return err
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
