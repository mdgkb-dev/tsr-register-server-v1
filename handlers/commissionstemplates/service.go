package commissionstemplates

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.CommissionTemplate) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	//item.SetIDForChildren()
	//err = CommissionTemplategroup.CreateService(s.helper).UpsertMany(item.CommissionTemplateGroups)
	if err != nil {
		return err
	}
	//err = CommissionTemplatediagnosis.CreateService(s.helper).CreateMany(item.CommissionTemplateDiagnosis)
	//if err != nil {
	//	return err
	//}
	return err
}

func (s *Service) GetAll() (models.CommissionsTemplates, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.CommissionTemplate, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.CommissionTemplate) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	//item.SetIDForChildren()

	//CommissionTemplateGroupService := CommissionTemplategroup.CreateService(s.helper)
	//err = CommissionTemplateGroupService.UpsertMany(item.CommissionTemplateGroups)
	//if err != nil {
	//	return err
	//}
	//err = CommissionTemplateGroupService.DeleteMany(item.CommissionsTemplatesForDelete)
	//if err != nil {
	//	return err
	//}

	//CommissionTemplateDiagnosisService := CommissionTemplatediagnosis.CreateService(s.helper)
	//err = CommissionTemplateDiagnosisService.UpsertMany(item.CommissionTemplateDiagnosis)
	//if err != nil {
	//	return err
	//}
	//err = CommissionTemplateDiagnosisService.DeleteMany(item.CommissionTemplateDiagnosisForDelete)
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
