package human

import (
	"mdgkb/tsr-tegister-server-v1/handlers/contact"
	"mdgkb/tsr-tegister-server-v1/handlers/document"
	"mdgkb/tsr-tegister-server-v1/handlers/fileInfo"
	"mdgkb/tsr-tegister-server-v1/handlers/insuranceCompanyToHuman"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) Create(item *models.Human) error {
	if item == nil {
		return nil
	}
	err := contact.CreateService(s.repository.getDB(), s.helper).Create(item.Contact)
	if err != nil {
		return err
	}
	item.ContactID = item.Contact.ID

	err = fileInfo.CreateService(s.repository.getDB(), s.helper).Create(item.Photo)
	if err != nil {
		return err
	}
	if item.Photo != nil {
		item.PhotoId.UUID = item.Photo.ID
	}

	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	err = insuranceCompanyToHuman.CreateService(s.repository.getDB(), s.helper).CreateMany(item.InsuranceCompanyToHuman)
	if err != nil {
		return err
	}
	err = document.CreateService(s.repository.getDB(), s.helper).CreateMany(item.Documents)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Human) error {
	if item == nil {
		return nil
	}
	err := contact.CreateService(s.repository.getDB(), s.helper).Upsert(item.Contact)
	if err != nil {
		return err
	}
	item.ContactID = item.Contact.ID

	err = fileInfo.CreateService(s.repository.getDB(), s.helper).Upsert(item.Photo)
	if err != nil {
		return err
	}
	if item.Photo != nil {
		item.PhotoId.UUID = item.Photo.ID
	}

	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	insuranceCompanyToHumanService := insuranceCompanyToHuman.CreateService(s.repository.getDB(), s.helper)
	err = insuranceCompanyToHumanService.UpsertMany(item.InsuranceCompanyToHuman)
	if err != nil {
		return err
	}
	err = insuranceCompanyToHumanService.DeleteMany(item.InsuranceCompanyToHumanForDelete)
	if err != nil {
		return err
	}
	documentService := document.CreateService(s.repository.getDB(), s.helper)
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

func (s *Service) Delete(id uuid.UUID) error {
	human, err := s.repository.get(id)
	if err != nil {
		return err
	}
	human.SetDeleteIdForChildren()
	err = contact.CreateService(s.repository.getDB(), s.helper).Delete(human.ContactID)
	if err != nil {
		return err
	}
	err = fileInfo.CreateService(s.repository.getDB(), s.helper).Delete(human.PhotoId)
	if err != nil {
		return err
	}
	err = insuranceCompanyToHuman.CreateService(s.repository.getDB(), s.helper).DeleteMany(human.InsuranceCompanyToHumanForDelete)
	if err != nil {
		return err
	}
	err = document.CreateService(s.repository.getDB(), s.helper).DeleteMany(human.DocumentsForDelete)
	if err != nil {
		return err
	}
	err = s.repository.delete(id)
	if err != nil {
		return err
	}
	return nil
}
