package patients

import (
	"mdgkb/tsr-tegister-server-v1/handlers/chestcircumference"
	"mdgkb/tsr-tegister-server-v1/handlers/chopscaletests"
	"mdgkb/tsr-tegister-server-v1/handlers/disability"
	"mdgkb/tsr-tegister-server-v1/handlers/headcircumference"
	"mdgkb/tsr-tegister-server-v1/handlers/heightweight"
	"mdgkb/tsr-tegister-server-v1/handlers/hmfsescaletests"
	"mdgkb/tsr-tegister-server-v1/handlers/human"
	"mdgkb/tsr-tegister-server-v1/handlers/patientdiagnosis"
	"mdgkb/tsr-tegister-server-v1/handlers/patientdrugregimen"
	"mdgkb/tsr-tegister-server-v1/handlers/registergroupstopatients"
	"mdgkb/tsr-tegister-server-v1/handlers/registertopatient"
	"mdgkb/tsr-tegister-server-v1/handlers/representativetopatient"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.Patient) error {
	err := human.CreateService(s.helper).Create(item.Human)
	if err != nil {
		return err
	}
	item.HumanID = item.Human.ID
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = representativetopatient.CreateService(s.helper).CreateMany(item.RepresentativeToPatient)
	if err != nil {
		return err
	}
	err = heightweight.CreateService(s.helper).CreateMany(item.HeightWeight)
	if err != nil {
		return err
	}
	err = chestcircumference.CreateService(s.helper).CreateMany(item.ChestCircumference)
	if err != nil {
		return err
	}
	err = headcircumference.CreateService(s.helper).CreateMany(item.HeadCircumference)
	if err != nil {
		return err
	}
	err = disability.CreateService(s.helper).CreateMany(item.Disabilities)
	if err != nil {
		return err
	}
	err = patientdiagnosis.CreateService(s.helper).CreateMany(item.PatientDiagnosis)
	if err != nil {
		return err
	}
	err = patientdrugregimen.CreateService(s.helper).CreateMany(item.PatientDrugRegimen)
	if err != nil {
		return err
	}
	err = registertopatient.CreateService(s.helper).CreateMany(item.RegisterToPatient)
	if err != nil {
		return err
	}
	err = chopscaletests.CreateService(s.helper).CreateMany(item.ChopScaleTests)
	if err != nil {
		return err
	}
	err = hmfsescaletests.CreateService(s.helper).CreateMany(item.HmfseScaleTests)
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
	err := human.CreateService(s.helper).Update(item.Human)
	if err != nil {
		return err
	}
	item.HumanID = item.Human.ID
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	representativeToPatientService := representativetopatient.CreateService(s.helper)
	err = representativeToPatientService.UpsertMany(item.RepresentativeToPatient)
	if err != nil {
		return err
	}
	err = representativeToPatientService.DeleteMany(item.RepresentativeToPatientForDelete)
	if err != nil {
		return err
	}
	heightWeightService := heightweight.CreateService(s.helper)
	err = heightWeightService.UpsertMany(item.HeightWeight)
	if err != nil {
		return err
	}
	err = heightWeightService.DeleteMany(item.HeightWeightForDelete)
	if err != nil {
		return err
	}
	chestCircumferenceService := chestcircumference.CreateService(s.helper)
	err = chestCircumferenceService.UpsertMany(item.ChestCircumference)
	if err != nil {
		return err
	}
	err = chestCircumferenceService.DeleteMany(item.ChestCircumferenceForDelete)
	if err != nil {
		return err
	}
	headCircumferenceService := headcircumference.CreateService(s.helper)
	err = headCircumferenceService.UpsertMany(item.HeadCircumference)
	if err != nil {
		return err
	}
	err = headCircumferenceService.DeleteMany(item.HeadCircumferenceForDelete)
	if err != nil {
		return err
	}
	disabilityService := disability.CreateService(s.helper)
	err = disabilityService.UpsertMany(item.Disabilities)
	if err != nil {
		return err
	}
	err = disabilityService.DeleteMany(item.DisabilitiesForDelete)
	if err != nil {
		return err
	}

	patientDiagnosisService := patientdiagnosis.CreateService(s.helper)
	err = patientDiagnosisService.UpsertMany(item.PatientDiagnosis)
	if err != nil {
		return err
	}
	err = patientDiagnosisService.DeleteMany(item.PatientDiagnosisForDelete)
	if err != nil {
		return err
	}
	patientDrugRegimenService := patientdrugregimen.CreateService(s.helper)
	err = patientDrugRegimenService.UpsertMany(item.PatientDrugRegimen)
	if err != nil {
		return err
	}
	err = patientDrugRegimenService.DeleteMany(item.PatientDrugRegimenForDelete)
	if err != nil {
		return err
	}
	registerToPatientService := registertopatient.CreateService(s.helper)
	err = registerToPatientService.UpsertMany(item.RegisterToPatient)
	if err != nil {
		return err
	}
	err = registerToPatientService.DeleteMany(item.RegisterToPatientForDelete)
	if err != nil {
		return err
	}

	registerGroupsToPatientsService := registergroupstopatients.CreateService(s.helper)
	err = registerGroupsToPatientsService.UpsertMany(item.RegisterGroupsToPatient)
	if err != nil {
		return err
	}
	err = registerGroupsToPatientsService.DeleteMany(item.RegisterGroupsToPatientsForDelete)
	if err != nil {
		return err
	}
	chopScaleTestService := chopscaletests.CreateService(s.helper)
	err = chopScaleTestService.UpsertMany(item.ChopScaleTests)
	if err != nil {
		return err
	}
	err = chopScaleTestService.DeleteMany(item.ChopScaleTestsForDelete)
	if err != nil {
		return err
	}
	hmfseScaleTestService := hmfsescaletests.CreateService(s.helper)
	err = hmfseScaleTestService.UpsertMany(item.HmfseScaleTests)
	if err != nil {
		return err
	}
	err = hmfseScaleTestService.DeleteMany(item.HmfseScaleTestsForDelete)
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
	patient.SetDeleteIDForChildren()
	err = human.CreateService(s.helper).Delete(patient.HumanID)
	if err != nil {
		return err
	}
	err = representativetopatient.CreateService(s.helper).DeleteMany(patient.RepresentativeToPatientForDelete)
	if err != nil {
		return err
	}
	err = heightweight.CreateService(s.helper).DeleteMany(patient.HeightWeightForDelete)
	if err != nil {
		return err
	}
	err = chestcircumference.CreateService(s.helper).DeleteMany(patient.ChestCircumferenceForDelete)
	if err != nil {
		return err
	}
	err = headcircumference.CreateService(s.helper).DeleteMany(patient.HeadCircumferenceForDelete)
	if err != nil {
		return err
	}
	err = disability.CreateService(s.helper).DeleteMany(patient.DisabilitiesForDelete)
	if err != nil {
		return err
	}
	err = patientdiagnosis.CreateService(s.helper).DeleteMany(patient.PatientDiagnosisForDelete)
	if err != nil {
		return err
	}
	err = patientdrugregimen.CreateService(s.helper).DeleteMany(patient.PatientDrugRegimenForDelete)
	if err != nil {
		return err
	}
	err = registertopatient.CreateService(s.helper).DeleteMany(patient.RegisterToPatientForDelete)
	if err != nil {
		return err
	}
	//err = registerPropertyToPatient.CreateService(s.helper).DeleteMany(patient.RegisterPropertyToPatientForDelete)
	//if err != nil {
	//	return err
	//}
	//err = registerPropertySetToPatient.CreateService(s.helper).DeleteMany(patient.RegisterPropertySetToPatientForDelete)
	//if err != nil {
	//	return err
	//}
	err = s.repository.delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetBySearch(query *string) ([]*models.Patient, error) {
	queryRu := s.helper.Util.TranslitToRu(*query)
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
