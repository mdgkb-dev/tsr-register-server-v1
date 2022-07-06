package patientDiagnosis

import (
	"fmt"
	"mdgkb/tsr-tegister-server-v1/handlers/patientDiagnosisAnamnesis"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items []*models.PatientDiagnosis) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	err = patientDiagnosisAnamnesis.CreateService(s.helper).CreateMany(models.GetPatientDiagnosisAnamnesis(items))
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
	patientDiagnosisAnamnesisService := patientDiagnosisAnamnesis.CreateService(s.helper)
	err = patientDiagnosisAnamnesisService.UpsertMany(models.GetPatientDiagnosisAnamnesis(items))
	if err != nil {
		return err
	}

	err = patientDiagnosisAnamnesisService.DeleteMany(models.GetPatientDiagnosisAnamnesisForDelete(items))
	if err != nil {
		return err
	}
	fmt.Println(items)
	return s.repository.upsertMany(items)
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
