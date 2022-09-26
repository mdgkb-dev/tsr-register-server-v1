package registergroupstopatients

import (
	"mdgkb/tsr-tegister-server-v1/handlers/registerpropertyotherstopatient"
	"mdgkb/tsr-tegister-server-v1/handlers/registerpropertysettopatient"
	"mdgkb/tsr-tegister-server-v1/handlers/registerpropertytopatient"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) UpsertMany(items models.RegisterGroupsToPatients) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	registerPropertyToPatientService := registerpropertytopatient.CreateService(s.helper)
	err = registerPropertyToPatientService.UpsertMany(items.GetRegisterPropertiesToPatients())
	if err != nil {
		return err
	}
	err = registerPropertyToPatientService.DeleteMany(items.GetRegisterPropertiesToPatientsForDelete())
	if err != nil {
		return err
	}
	registerPropertySetToPatientService := registerpropertysettopatient.CreateService(s.helper)
	err = registerPropertySetToPatientService.UpsertMany(items.GetRegisterPropertySetToPatient())
	if err != nil {
		return err
	}
	err = registerPropertySetToPatientService.DeleteMany(items.GetRegisterPropertySetToPatientForDelete())
	if err != nil {
		return err
	}

	registerPropertyOthersToPatientService := registerpropertyotherstopatient.CreateService(s.helper)
	err = registerPropertyOthersToPatientService.UpsertMany(items.GetRegisterPropertyOthersToPatient())
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
