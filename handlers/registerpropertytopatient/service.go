package registerpropertytopatient

import (
	"mdgkb/tsr-tegister-server-v1/handlers/registerpropertiestopatientstofileinfos"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items []*models.Answer) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.createMany(items)
}

func (s *Service) UpsertMany(items models.Answers) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	registerPropertiesToPatientsToFileInfosService := registerpropertiestopatientstofileinfos.CreateService(s.helper)
	err = registerPropertiesToPatientsToFileInfosService.UpsertMany(items.GetRegisterPropertiesToPatientsToFileInfos())
	if err != nil {
		return err
	}
	err = registerPropertiesToPatientsToFileInfosService.DeleteMany(items.GetRegisterPropertiesToPatientsToFileInfosForDelete())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
