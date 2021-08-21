package human

import (
	"mdgkb/tsr-tegister-server-v1/handlers/document"
	"mdgkb/tsr-tegister-server-v1/handlers/insuranceCompanyToHuman"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.Human) error {
	if item == nil {
		return nil
	}
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	err = insuranceCompanyToHuman.CreateService(s.repository.getDB()).CreateMany(item.InsuranceCompanyToHuman)
	if err != nil {
		return err
	}
	err = document.CreateService(s.repository.getDB()).CreateMany(item.Documents)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Human) error {
	if item == nil {
		return nil
	}
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	insuranceCompanyToHumanService := insuranceCompanyToHuman.CreateService(s.repository.getDB())
	err = insuranceCompanyToHumanService.UpsertMany(item.InsuranceCompanyToHuman)
	if err != nil {
		return err
	}
	err = insuranceCompanyToHumanService.DeleteMany(item.InsuranceCompanyToHumanForDelete)
	if err != nil {
		return err
	}
	documentService := document.CreateService(s.repository.getDB())
	err = documentService.UpsertMany(item.Documents)
	if err != nil {
		return err
	}
	err = documentService.DeleteMany(item.DocumentsForDelete)
	if err != nil {
		return err
	}
	return nil
}
