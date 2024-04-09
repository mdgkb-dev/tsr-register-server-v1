package researches

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/handlers/patients"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(c context.Context, item *models.Research) error {
	err := R.Create(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	// err = Researchgroup.CreateService(s.helper).UpsertMany(item.ResearchGroups)
	if err != nil {
		return err
	}
	//err = Researchdiagnosis.CreateService(s.helper).CreateMany(item.ResearchDiagnosis)
	//if err != nil {
	//	return err
	//}
	return err
}

func (s *Service) GetAll(c context.Context) (models.ResearchesWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.Research, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(c context.Context, item *models.Research) error {
	err := R.Update(c, item)
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

func (s *Service) Delete(c context.Context, id *string) error {
	return R.Delete(c, id)
}

func (s *Service) GetResearchAndPatient(c context.Context, researchID string, patientID string) (*models.Research, *models.Patient, error) {
	research, err := R.Get(c, researchID)
	if err != nil {
		return nil, nil, err
	}

	patient, err := patients.S.Get(c, patientID)
	if err != nil {
		return nil, nil, err
	}
	return research, patient, nil
}
