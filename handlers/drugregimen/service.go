package drugregimen

import (
	"mdgkb/tsr-tegister-server-v1/handlers/drugregimenblock"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) CreateMany(items []*models.DrugRegimen) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	err = drugregimenblock.CreateService(s.helper).CreateMany(models.GetDrugRegimenBlocks(items))
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items []*models.DrugRegimen) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
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

func (s *Service) DeleteMany(idPool []string) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
