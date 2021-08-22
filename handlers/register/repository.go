package register

import (
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Register) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items []*models.Register, err error) {
	err = r.db.NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Register, error) {
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
		Relation("RegisterToPatient.Patient.Human").
		Relation("RegisterToPatient.Patient.RegisterPropertyToPatient.RegisterProperty").
		Relation("RegisterToPatient.Patient.RegisterPropertyToPatient.RegisterProperty").
		Relation("RegisterToPatient.Patient.RegisterPropertySetToPatient.RegisterPropertySet").
		Where("register.id = ?", *id).Scan(r.ctx)
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
