package patients

import (
	"mdgkb/tsr-tegister-server-v1/handlers/human"
	"mdgkb/tsr-tegister-server-v1/handlers/humans"
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
	err := humans.CreateService(s.helper).Update(item.Human)
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

	patientDrugRegimenService := patientdrugregimen.CreateService(s.helper)
	err = patientDrugRegimenService.UpsertMany(item.PatientDrugRegimen)
	if err != nil {
		return err
	}
	err = patientDrugRegimenService.DeleteMany(item.PatientDrugRegimenForDelete)
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
