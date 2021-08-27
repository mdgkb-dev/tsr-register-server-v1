package drugRegimen

import (
	"mdgkb/tsr-tegister-server-v1/handlers/drugRegimenBlock"
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
	err = drugRegimenBlock.CreateService(s.repository.getDB()).CreateMany(models.GetDrugRegimenBlocks(items))
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
	drugRegimenBlockService := drugRegimenBlock.CreateService(s.repository.getDB())
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
