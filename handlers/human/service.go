package human

import (
	"mdgkb/tsr-tegister-server-v1/handlers/contact"
	"mdgkb/tsr-tegister-server-v1/handlers/fileinfos"
	"mdgkb/tsr-tegister-server-v1/handlers/insurancecompanytohuman"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) Create(item *models.Human) error {
	if item == nil {
		return nil
	}
	err := contact.CreateService(s.helper).Create(item.Contact)
	if err != nil {
		return err
	}
	item.ContactID = item.Contact.ID

	err = fileinfos.CreateService(s.helper).Create(item.Photo)
	if err != nil {
		return err
	}
	if item.Photo != nil {
		item.PhotoID = item.Photo.ID
	}

	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	err = insurancecompanytohuman.CreateService(s.helper).CreateMany(item.InsuranceCompanyToHuman)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Human) error {
	if item == nil {
		return nil
	}
	err := contact.CreateService(s.helper).Upsert(item.Contact)
	if err != nil {
		return err
	}
	item.ContactID = item.Contact.ID

	err = fileinfos.CreateService(s.helper).Upsert(item.Photo)
	if err != nil {
		return err
	}
	if item.Photo != nil {
		item.PhotoID = item.Photo.ID
	}

	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	insuranceCompanyToHumanService := insurancecompanytohuman.CreateService(s.helper)
	err = insuranceCompanyToHumanService.UpsertMany(item.InsuranceCompanyToHuman)
	if err != nil {
		return err
	}
	err = insuranceCompanyToHumanService.DeleteMany(item.InsuranceCompanyToHumanForDelete)
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
	human.SetDeleteIDForChildren()
	err = contact.CreateService(s.helper).Delete(human.ContactID)
	if err != nil {
		return err
	}
	err = fileinfos.CreateService(s.helper).Delete(human.PhotoID)
	if err != nil {
		return err
	}
	err = insurancecompanytohuman.CreateService(s.helper).DeleteMany(human.InsuranceCompanyToHumanForDelete)
	if err != nil {
		return err
	}
	return nil
}
