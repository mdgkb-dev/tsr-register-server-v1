package patientsdomains

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/pro-assistance/pro-assister/middleware"

	"github.com/google/uuid"
)

func (s *Service) Create(c context.Context, item *models.PatientDomain) error {
	err := R.Create(c, item)
	if err != nil {
		return err
	}
	return nil
}

func createFromPatientID(c context.Context, patientID uuid.NullUUID) (models.PatientsDomains, error) {
	domainsIDS := middleware.ClaimDomainIDS.FromContextSlice(c)
	items := make(models.PatientsDomains, len(domainsIDS))
	for i := range domainsIDS {
		dID, err := uuid.Parse(domainsIDS[i])
		if err != nil {
			return nil, err
		}
		d := models.PatientDomain{PatientID: patientID, DomainID: uuid.NullUUID{UUID: dID, Valid: true}}
		items[i] = &d
	}
	return items, nil
}

func (s *Service) AddToDomain(c context.Context, patientID uuid.NullUUID) error {
	items, err := createFromPatientID(c, patientID)
	if err != nil {
		return err
	}
	err = R.AddToDomain(c, items)
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

func (s *Service) PatientInDomain(c context.Context, patientID string) (bool, error) {
	return R.PatientInDomain(c, patientID)
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}
