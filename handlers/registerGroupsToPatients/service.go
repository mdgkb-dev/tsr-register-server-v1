package registerGroupsToPatients

import (
	"fmt"
	"mdgkb/tsr-tegister-server-v1/handlers/registerPropertyOthersToPatient"
	"mdgkb/tsr-tegister-server-v1/handlers/registerPropertySetToPatient"
	"mdgkb/tsr-tegister-server-v1/handlers/registerPropertyToPatient"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) UpsertMany(items models.RegisterGroupsToPatients) error {
	fmt.Println("items")
	fmt.Println("items")
	fmt.Println("items")
	fmt.Println("items")
	fmt.Println("items")
	fmt.Println("items")
	fmt.Println(items)
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	registerPropertyToPatientService := registerPropertyToPatient.CreateService(s.helper)
	err = registerPropertyToPatientService.UpsertMany(items.GetRegisterPropertiesToPatients())
	if err != nil {
		return err
	}
	err = registerPropertyToPatientService.DeleteMany(items.GetRegisterPropertiesToPatientsForDelete())
	if err != nil {
		return err
	}
	registerPropertySetToPatientService := registerPropertySetToPatient.CreateService(s.helper)
	err = registerPropertySetToPatientService.UpsertMany(items.GetRegisterPropertySetToPatient())
	if err != nil {
		return err
	}
	err = registerPropertySetToPatientService.DeleteMany(items.GetRegisterPropertySetToPatientForDelete())
	if err != nil {
		return err

	}

	registerPropertyOthersToPatientService := registerPropertyOthersToPatient.CreateService(s.helper)
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
