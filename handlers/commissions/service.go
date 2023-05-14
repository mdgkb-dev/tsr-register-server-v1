package commissions

import (
	"mdgkb/tsr-tegister-server-v1/handlers/commissionsdoctors"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.Commission) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = commissionsdoctors.CreateService(s.helper).UpsertMany(item.CommissionsDoctors)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) GetAll() (models.CommissionsWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.Commission, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.Commission) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	//item.SetIDForChildren()

	//CommissionGroupService := Commissiongroup.CreateService(s.helper)
	//err = CommissionGroupService.UpsertMany(item.CommissionGroups)
	//if err != nil {
	//	return err
	//}
	//err = CommissionGroupService.DeleteMany(item.CommissionsForDelete)
	//if err != nil {
	//	return err
	//}

	//CommissionDiagnosisService := Commissiondiagnosis.CreateService(s.helper)
	//err = CommissionDiagnosisService.UpsertMany(item.CommissionDiagnosis)
	//if err != nil {
	//	return err
	//}
	//err = CommissionDiagnosisService.DeleteMany(item.CommissionDiagnosisForDelete)
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
