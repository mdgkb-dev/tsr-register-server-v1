package patientsdiagnosis

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) Create(c context.Context, item *models.PatientDiagnosis) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (item models.PatientDiagnosisWithCount, err error) {
	item.PatientDiagnosis = make([]*models.PatientDiagnosis, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&item.PatientDiagnosis)

	r.helper.SQL.ExtractQueryFilter(c).HandleQuery(query)
	item.Count, err = query.ScanAndCount(c)
	return item, err
}

func (r *Repository) Get(c context.Context, slug string) (*models.PatientDiagnosis, error) {
	item := models.PatientDiagnosis{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Where("?TableAlias.id = ?", slug).
		Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.PatientDiagnosis{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.PatientDiagnosis) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}
