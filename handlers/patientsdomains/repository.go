package patientsdomains

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) Create(c context.Context, item *models.PatientDomain) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll(c context.Context) (item models.PatientsDomainsWithCount, err error) {
	item.PatientsDomains = make(models.PatientsDomains, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&item.PatientsDomains)

	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) Get(c context.Context, slug string) (*models.PatientDomain, error) {
	item := models.PatientDomain{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Where("?TableAlias.id = ?", slug).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) PatientInDomain(c context.Context, patientID string, domainID string) (bool, error) {
	return r.helper.DB.IDB(c).NewSelect().Model((*models.PatientDomain)(nil)).
		Where("?TableAlias.patient_id = ?", patientID).
		Where("?TableAlias.domain_id = ?", domainID).
		Exists(r.ctx)
}

func (r *Repository) AddToDomain(c context.Context, item *models.PatientDomain) error {
	_, err := r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.PatientDomain{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(c context.Context, item *models.PatientDomain) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
