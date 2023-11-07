package representativesdomains

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/middleware"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) Create(c context.Context, item *models.RepresentativeDomain) error {
	err := R.Create(c, item)
	if err != nil {
		return err
	}
	return nil
}

func createFromRepresentativeID(c context.Context, representativeID uuid.NullUUID) (models.RepresentativesDomains, error) {
	domainsIDS := middleware.ClaimDomainIDS.FromContextSlice(c)
	items := make(models.RepresentativesDomains, 0)
	for i := range domainsIDS {
		dID, err := uuid.Parse(domainsIDS[i])
		if err != nil {
			return nil, err
		}
		d := models.RepresentativeDomain{RepresentativeID: representativeID, DomainID: uuid.NullUUID{UUID: dID, Valid: true}}
		items = append(items, &d)
	}
	return items, nil
}

func (s *Service) AddToDomain(c context.Context, representativeID uuid.NullUUID) error {
	items, err := createFromRepresentativeID(c, representativeID)
	if err != nil {
		return err
	}
	err = R.AddToDomain(c, items)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(c context.Context, item *models.RepresentativeDomain) error {
	err := R.Update(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(c context.Context) (models.RepresentativesDomainsWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, slug string) (*models.RepresentativeDomain, error) {
	item, err := R.Get(c, slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) RepresentativeInDomain(c context.Context, RepresentativeID string) (bool, error) {
	return R.RepresentativeInDomain(c, RepresentativeID)
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}
