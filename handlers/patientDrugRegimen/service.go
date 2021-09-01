package patientDrugRegimen

import (
	"mdgkb/tsr-tegister-server-v1/handlers/patientDrugRegimenItem"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) CreateMany(items []*models.PatientDrugRegimen) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	err = patientDrugRegimenItem.CreateService(s.repository.getDB()).CreateMany(models.GetPatientDrugRegimenItems(items))
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items []*models.PatientDrugRegimen) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	err = patientDrugRegimenItem.CreateService(s.repository.getDB()).UpsertMany(models.GetPatientDrugRegimenItems(items))
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteMany(idPool []string) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
