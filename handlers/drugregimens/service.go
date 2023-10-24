package drugregimens

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/handlers/drugregimenblock"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(c context.Context, items models.DrugRegimens) error {
	if len(items) == 0 {
		return nil
	}
	err := R.CreateMany(c, items)
	if err != nil {
		return err
	}
	err = drugregimenblock.CreateService(s.helper).CreateMany(models.GetDrugRegimenBlocks(items))
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(c context.Context, items models.DrugRegimens) error {
	if len(items) == 0 {
		return nil
	}
	err := R.UpsertMany(c, items)
	if err != nil {
		return err
	}
	drugRegimenBlockService := drugregimenblock.CreateService(s.helper)
	err = drugRegimenBlockService.UpsertMany(models.GetDrugRegimenBlocks(items))
	if err != nil {
		return err
	}

	err = drugRegimenBlockService.DeleteMany(models.GetDrugRegimenBlocksForDelete(items))
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteMany(c context.Context, idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return R.DeleteMany(c, idPool)
}

func (s *Service) Create(c context.Context, item *models.DrugRegimen) error {
	err := s.repository.Create(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(c context.Context, item *models.DrugRegimen) error {
	err := s.repository.Update(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(c context.Context) (models.DrugRegimensWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, slug string) (*models.DrugRegimen, error) {
	item, err := s.repository.Get(c, slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}

func (r *Service) GetByParameters(c context.Context, drugDozeID uuid.NullUUID, months uint, weight uint) (*models.DrugRegimen, error) {
	return R.GetByParameters(c, drugDozeID, months, weight)
}
