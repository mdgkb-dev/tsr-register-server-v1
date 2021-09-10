package register

import (
	"mdgkb/tsr-tegister-server-v1/helpers/httpHelper"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Register) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items []*models.Register, err error) {
	err = r.db.NewSelect().Model(&items).Relation("RegisterDiagnosis").Scan(r.ctx)
	return items, err
}

func (r *Repository) get(queryFilter *httpHelper.QueryFilter) (*models.Register, error) {
	item := models.Register{}
	err := r.db.NewSelect().
		Model(&item).
		Relation("RegisterGroupToRegister.RegisterGroup").
		Relation("RegisterDiagnosis.MkbDiagnosis.MkbSubDiagnosis").
		Relation("RegisterDiagnosis.MkbDiagnosis.MkbGroup").
		Relation("RegisterDiagnosis.MkbSubDiagnosis").
		Relation("RegisterGroupToRegister.RegisterGroup.RegisterPropertyToRegisterGroup.RegisterProperty.ValueType").
		Relation("RegisterGroupToRegister.RegisterGroup.RegisterPropertyToRegisterGroup.RegisterProperty.RegisterPropertySet").
		Relation("RegisterGroupToRegister.RegisterGroup.RegisterPropertyToRegisterGroup.RegisterProperty.RegisterPropertyRadio").
		Relation("RegisterToPatient.Patient.Human", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("patient__human.surname", "patient__human.name", "patient__human.patronymic").Offset(*queryFilter.Pagination.Offset).
				Limit(*queryFilter.Pagination.Limit)
		}).
		Relation("RegisterToPatient.Patient.RegisterPropertyToPatient.RegisterProperty").
		Relation("RegisterToPatient.Patient.RegisterPropertyToPatient.RegisterProperty").
		Relation("RegisterToPatient.Patient.RegisterPropertySetToPatient.RegisterPropertySet").
		Where("register.id = ?", *queryFilter.ID).Scan(r.ctx)
	if err != nil {
		return nil, err
	}
	item.RegisterToPatientCount, err = r.db.NewSelect().Model((*models.RegisterToPatient)(nil)).Where("register_id = ?", *queryFilter.ID).Count(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Register{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Register) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
