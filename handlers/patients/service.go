package patients

import (
	"mdgkb/tsr-tegister-server-v1/handlers/chestcircumference"
	"mdgkb/tsr-tegister-server-v1/handlers/chopscaletests"
	"mdgkb/tsr-tegister-server-v1/handlers/headcircumference"
	"mdgkb/tsr-tegister-server-v1/handlers/heightweight"
	"mdgkb/tsr-tegister-server-v1/handlers/hmfsescaletests"
	"mdgkb/tsr-tegister-server-v1/handlers/human"
	"mdgkb/tsr-tegister-server-v1/handlers/patientdrugregimen"
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
	err = s.repository.Create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = representativetopatient.CreateService(s.helper).CreateMany(item.PatientsRepresentatives)
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
	err = patientdrugregimen.CreateService(s.helper).CreateMany(item.PatientDrugRegimen)
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
	return s.repository.GetAll()
}

func (s *Service) Get(id string) (*models.Patient, error) {
	item, err := s.repository.Get(id)
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
	err = s.repository.Update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	representativeToPatientService := representativetopatient.CreateService(s.helper)
	err = representativeToPatientService.UpsertMany(item.PatientsRepresentatives)
	if err != nil {
		return err
	}
	err = representativeToPatientService.DeleteMany(item.PatientsRepresentativesForDelete)
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

	patientDrugRegimenService := patientdrugregimen.CreateService(s.helper)
	err = patientDrugRegimenService.UpsertMany(item.PatientDrugRegimen)
	if err != nil {
		return err
	}
	err = patientDrugRegimenService.DeleteMany(item.PatientDrugRegimenForDelete)
	if err != nil {
		return err
	}

	//registerGroupsToPatientsService := patientresearchsections.CreateService(s.helper)
	//err = registerGroupsToPatientsService.UpsertMany(item.RegisterGroupsToPatient)
	//if err != nil {
	//	return err
	//}
	//err = registerGroupsToPatientsService.DeleteMany(item.RegisterGroupsToPatientsForDelete)
	//if err != nil {
	//	return err
	//}
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

func (s *Service) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SetQueryFilter(c *gin.Context) (err error) {
	err = s.repository.SetQueryFilter(c)
	return err
}
