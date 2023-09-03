package researches

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/handlers/patients"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.Research) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	//err = Researchgroup.CreateService(s.helper).UpsertMany(item.ResearchGroups)
	if err != nil {
		return err
	}
	//err = Researchdiagnosis.CreateService(s.helper).CreateMany(item.ResearchDiagnosis)
	//if err != nil {
	//	return err
	//}
	return err
}

func (s *Service) GetAll(c context.Context) (models.Researches, error) {
	items, err := s.repository.getAll(c)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.Research, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.Research) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	//ResearchGroupService := Researchgroup.CreateService(s.helper)
	//err = ResearchGroupService.UpsertMany(item.ResearchGroups)
	//if err != nil {
	//	return err
	//}
	//err = ResearchGroupService.DeleteMany(item.ResearchesForDelete)
	//if err != nil {
	//	return err
	//}

	//ResearchDiagnosisService := Researchdiagnosis.CreateService(s.helper)
	//err = ResearchDiagnosisService.UpsertMany(item.ResearchDiagnosis)
	//if err != nil {
	//	return err
	//}
	//err = ResearchDiagnosisService.DeleteMany(item.ResearchDiagnosisForDelete)
	//if err != nil {
	//	return err
	//}
	return err
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) GetValueTypes() (models.ValueTypes, error) {
	items, err := s.repository.getValueTypes()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}

func (s *Service) GetResearchAndPatient(ctx context.Context, researchId string, patientId string) (*models.Research, *models.Patient, error) {
	research, err := R.get(researchId)
	if err != nil {
		return nil, nil, err
	}

	patient, err := patients.S.Get(ctx, patientId)
	if err != nil {
		return nil, nil, err
	}
	return research, patient, nil
}
