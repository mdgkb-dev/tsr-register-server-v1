package patientsdomains

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(c context.Context, item *models.PatientDomain) error {
	err := R.Create(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) AddToDomain(c context.Context, item *models.PatientDomain) error {
	err := R.AddToDomain(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(c context.Context, item *models.PatientDomain) error {
	err := R.Update(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(c context.Context) (models.PatientsDomainsWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, slug string) (*models.PatientDomain, error) {
	item, err := R.Get(c, slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) PatientInDomain(c context.Context, patientID string, domainID string) (bool, error) {
	return R.PatientInDomain(c, patientID, domainID)
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}
