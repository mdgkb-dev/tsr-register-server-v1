package edv

import (
	"mdgkb/tsr-tegister-server-v1/handlers/fileinfos"
	"mdgkb/tsr-tegister-server-v1/handlers/period"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) CreateMany(items []*models.Edv) error {
	if len(items) == 0 {
		return nil
	}
	err := period.CreateService(s.helper).CreateMany(models.GetPeriodsFromEdv(items))
	if err != nil {
		return err
	}
	models.SetPeriodIDToEdv(items)
	err = fileinfos.CreateService(s.helper).CreateMany(models.GetFilesFromEdv(items))
	if err != nil {
		return err
	}
	models.SetFileInfoIDToEdv(items)
	return s.repository.createMany(items)
}

func (s *Service) UpsertMany(items []*models.Edv) error {
	if len(items) == 0 {
		return nil
	}
	err := period.CreateService(s.helper).UpsertMany(models.GetPeriodsFromEdv(items))
	if err != nil {
		return err
	}
	models.SetPeriodIDToEdv(items)
	err = fileinfos.CreateService(s.helper).UpsertMany(models.GetFilesFromEdv(items))
	if err != nil {
		return err
	}
	models.SetFileInfoIDToEdv(items)
	return s.repository.upsertMany(items)
}

func (s *Service) DeleteMany(idPool []string) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
