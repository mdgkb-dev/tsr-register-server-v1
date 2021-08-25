package patientDiagnosis

import (
	"mdgkb/tsr-tegister-server-v1/handlers/patientDiagnosisAnamnesis"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) CreateMany(items []*models.PatientDiagnosis) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	err = patientDiagnosisAnamnesis.CreateService(s.repository.getDB()).CreateMany(models.GetPatientDiagnosisAnamnesis(items))
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items []*models.PatientDiagnosis) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	patientDiagnosisAnamnesisService := patientDiagnosisAnamnesis.CreateService(s.repository.getDB())
	err = patientDiagnosisAnamnesisService.UpsertMany(models.GetPatientDiagnosisAnamnesis(items))
	if err != nil {
		return err
	}

	err = patientDiagnosisAnamnesisService.DeleteMany(models.GetPatientDiagnosisAnamnesisForDelete(items))
	if err != nil {
		return err
	}
	return s.repository.upsertMany(items)
}

func (s *Service) DeleteMany(idPool []string) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
