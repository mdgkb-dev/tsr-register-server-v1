package drugregimenblock

import (
	"mdgkb/tsr-tegister-server-v1/handlers/drugregimenblockitem"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) CreateMany(items []*models.DrugRegimenBlock) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	err = drugregimenblockitem.CreateService(s.helper).CreateMany(models.GetDrugRegimenBlockItems(items))
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items []*models.DrugRegimenBlock) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	drugRegimenBlockItemService := drugregimenblockitem.CreateService(s.helper)
	err = drugRegimenBlockItemService.UpsertMany(models.GetDrugRegimenBlockItems(items))
	if err != nil {
		return err
	}
	err = drugRegimenBlockItemService.DeleteMany(models.GetDrugRegimenBlockItemsForDelete(items))
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteMany(idPool []string) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
