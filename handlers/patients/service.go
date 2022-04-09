package patients

import (
	"github.com/gin-gonic/gin"
	"mdgkb/tsr-tegister-server-v1/handlers/chestCircumference"
	"mdgkb/tsr-tegister-server-v1/handlers/disability"
	"mdgkb/tsr-tegister-server-v1/handlers/headCircumference"
	"mdgkb/tsr-tegister-server-v1/handlers/heightWeight"
	"mdgkb/tsr-tegister-server-v1/handlers/human"
	"mdgkb/tsr-tegister-server-v1/handlers/patientDiagnosis"
	"mdgkb/tsr-tegister-server-v1/handlers/patientDrugRegimen"
	"mdgkb/tsr-tegister-server-v1/handlers/registerPropertyOthersToPatient"
	"mdgkb/tsr-tegister-server-v1/handlers/registerPropertySetToPatient"
	"mdgkb/tsr-tegister-server-v1/handlers/registerPropertyToPatient"
	"mdgkb/tsr-tegister-server-v1/handlers/registerToPatient"
	"mdgkb/tsr-tegister-server-v1/handlers/representativeToPatient"
	"mdgkb/tsr-tegister-server-v1/helpers"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.Patient) error {
	err := human.CreateService(s.repository.getDB()).Create(item.Human)
	if err != nil {
		return err
	}
	item.HumanID = item.Human.ID
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = representativeToPatient.CreateService(s.repository.getDB()).CreateMany(item.RepresentativeToPatient)
	if err != nil {
		return err
	}
	err = heightWeight.CreateService(s.repository.getDB()).CreateMany(item.HeightWeight)
	if err != nil {
		return err
	}
	err = chestCircumference.CreateService(s.repository.getDB()).CreateMany(item.ChestCircumference)
	if err != nil {
		return err
	}
	err = headCircumference.CreateService(s.repository.getDB()).CreateMany(item.HeadCircumference)
	if err != nil {
		return err
	}
	err = disability.CreateService(s.repository.getDB()).CreateMany(item.Disabilities)
	if err != nil {
		return err
	}
	err = patientDiagnosis.CreateService(s.repository.getDB()).CreateMany(item.PatientDiagnosis)
	if err != nil {
		return err
	}
	err = patientDrugRegimen.CreateService(s.repository.getDB()).CreateMany(item.PatientDrugRegimen)
	if err != nil {
		return err
	}
	err = registerToPatient.CreateService(s.repository.getDB()).CreateMany(item.RegisterToPatient)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) GetAll() (models.PatientsWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) GetOnlyNames() (models.PatientsWithCount, error) {
	return s.repository.getOnlyNames()
}

func (s *Service) Get(id *string, withDeleted bool) (*models.Patient, error) {
	item, err := s.repository.get(id, withDeleted)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.Patient) error {
	err := human.CreateService(s.repository.getDB()).Update(item.Human)
	if err != nil {
		return err
	}
	item.HumanID = item.Human.ID
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	representativeToPatientService := representativeToPatient.CreateService(s.repository.getDB())
	err = representativeToPatientService.UpsertMany(item.RepresentativeToPatient)
	if err != nil {
		return err
	}
	err = representativeToPatientService.DeleteMany(item.RepresentativeToPatientForDelete)
	if err != nil {
		return err
	}
	heightWeightService := heightWeight.CreateService(s.repository.getDB())
	err = heightWeightService.UpsertMany(item.HeightWeight)
	if err != nil {
		return err
	}
	err = heightWeightService.DeleteMany(item.HeightWeightForDelete)
	if err != nil {
		return err
	}
	chestCircumferenceService := chestCircumference.CreateService(s.repository.getDB())
	err = chestCircumferenceService.UpsertMany(item.ChestCircumference)
	if err != nil {
		return err
	}
	err = chestCircumferenceService.DeleteMany(item.ChestCircumferenceForDelete)
	if err != nil {
		return err
	}
	headCircumferenceService := headCircumference.CreateService(s.repository.getDB())
	err = headCircumferenceService.UpsertMany(item.HeadCircumference)
	if err != nil {
		return err
	}
	err = headCircumferenceService.DeleteMany(item.HeadCircumferenceForDelete)
	if err != nil {
		return err
	}
	disabilityService := disability.CreateService(s.repository.getDB())
	err = disabilityService.UpsertMany(item.Disabilities)
	if err != nil {
		return err
	}
	err = disabilityService.DeleteMany(item.DisabilitiesForDelete)
	if err != nil {
		return err
	}

	patientDiagnosisService := patientDiagnosis.CreateService(s.repository.getDB())
	err = patientDiagnosisService.UpsertMany(item.PatientDiagnosis)
	if err != nil {
		return err
	}
	err = patientDiagnosisService.DeleteMany(item.PatientDiagnosisForDelete)
	if err != nil {
		return err
	}
	patientDrugRegimenService := patientDrugRegimen.CreateService(s.repository.getDB())
	err = patientDrugRegimenService.UpsertMany(item.PatientDrugRegimen)
	if err != nil {
		return err
	}
	err = patientDrugRegimenService.DeleteMany(item.PatientDrugRegimenForDelete)
	if err != nil {
		return err
	}
	registerToPatientService := registerToPatient.CreateService(s.repository.getDB())
	err = registerToPatientService.UpsertMany(item.RegisterToPatient)
	if err != nil {
		return err
	}
	err = registerToPatientService.DeleteMany(item.RegisterToPatientForDelete)
	if err != nil {
		return err
	}
	registerPropertyToPatientService := registerPropertyToPatient.CreateService(s.repository.getDB())
	err = registerPropertyToPatientService.UpsertMany(item.RegisterPropertyToPatient)
	if err != nil {
		return err
	}
	err = registerPropertyToPatientService.DeleteMany(item.RegisterPropertyToPatientForDelete)
	if err != nil {
		return err
	}
	registerPropertySetToPatientService := registerPropertySetToPatient.CreateService(s.repository.getDB())
	err = registerPropertySetToPatientService.UpsertMany(item.RegisterPropertySetToPatient)
	if err != nil {
		return err
	}
	err = registerPropertySetToPatientService.DeleteMany(item.RegisterPropertySetToPatientForDelete)
	if err != nil {
		return err

	}

	registerPropertyOthersToPatientService := registerPropertyOthersToPatient.CreateService(s.repository.getDB())
	err = registerPropertyOthersToPatientService.UpsertMany(item.RegisterPropertyOthersPatient)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(id *string) error {
	patient, err := s.repository.get(id, false)
	if err != nil {
		return err
	}
	patient.SetDeleteIdForChildren()
	err = human.CreateService(s.repository.getDB()).Delete(patient.HumanID)
	if err != nil {
		return err
	}
	err = representativeToPatient.CreateService(s.repository.getDB()).DeleteMany(patient.RepresentativeToPatientForDelete)
	if err != nil {
		return err
	}
	err = heightWeight.CreateService(s.repository.getDB()).DeleteMany(patient.HeightWeightForDelete)
	if err != nil {
		return err
	}
	err = chestCircumference.CreateService(s.repository.getDB()).DeleteMany(patient.ChestCircumferenceForDelete)
	if err != nil {
		return err
	}
	err = headCircumference.CreateService(s.repository.getDB()).DeleteMany(patient.HeadCircumferenceForDelete)
	if err != nil {
		return err
	}
	err = disability.CreateService(s.repository.getDB()).DeleteMany(patient.DisabilitiesForDelete)
	if err != nil {
		return err
	}
	err = patientDiagnosis.CreateService(s.repository.getDB()).DeleteMany(patient.PatientDiagnosisForDelete)
	if err != nil {
		return err
	}
	err = patientDrugRegimen.CreateService(s.repository.getDB()).DeleteMany(patient.PatientDrugRegimenForDelete)
	if err != nil {
		return err
	}
	err = registerToPatient.CreateService(s.repository.getDB()).DeleteMany(patient.RegisterToPatientForDelete)
	if err != nil {
		return err
	}
	err = registerPropertyToPatient.CreateService(s.repository.getDB()).DeleteMany(patient.RegisterPropertyToPatientForDelete)
	if err != nil {
		return err
	}
	err = registerPropertySetToPatient.CreateService(s.repository.getDB()).DeleteMany(patient.RegisterPropertySetToPatientForDelete)
	if err != nil {
		return err
	}
	err = s.repository.delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetBySearch(query *string) ([]*models.Patient, error) {
	queryRu := helpers.TranslitToRu(*query)
	items, err := s.repository.getBySearch(&queryRu)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) GetDisabilities() (models.PatientsWithCount, error) {
	items, err := s.repository.getDisabilities()
	if err != nil {
		return items, err
	}
	return items, nil
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
