package drug

import (
	"mdgkb/tsr-tegister-server-v1/handlers/drugregimen"
	"mdgkb/tsr-tegister-server-v1/handlers/drugsdiagnosis"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Service) Create(item *models.Drug) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = drugregimen.CreateService(s.helper).CreateMany(item.DrugRegimens)
	if err != nil {
		return err
	}
	err = drugsdiagnosis.CreateService(s.helper).CreateMany(item.DrugsDiagnosis)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(diagnosis []uuid.UUID) ([]*models.Drug, error) {
	items, err := s.repository.getAll(diagnosis)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.Drug, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.Drug) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	drugRegimenService := drugregimen.CreateService(s.helper)
	err = drugRegimenService.UpsertMany(item.DrugRegimens)
	if err != nil {
		return err
	}
	err = drugRegimenService.DeleteMany(item.DrugRegimensForDelete)
	if err != nil {
		return err
	}

	drugsDiagnosisService := drugsdiagnosis.CreateService(s.helper)
	err = drugsDiagnosisService.UpsertMany(item.DrugsDiagnosis)
	if err != nil {
		return err
	}
	err = drugsDiagnosisService.DeleteMany(item.DrugsDiagnosisForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
