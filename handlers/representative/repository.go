package representative

import (
	"mdgkb/tsr-tegister-server-v1/helpers/httpHelper"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Representative) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll(pagination *httpHelper.Pagination) (items []*models.Representative, err error) {
	err = r.db.NewSelect().
		Model(&items).
		Relation("Human").
		Relation("RepresentativeToPatient.Patient.Human").
		Relation("RepresentativeToPatient.RepresentativeType").
		Offset(*pagination.Offset).
		Limit(*pagination.Limit).
		Order("human.surname").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) getOnlyNames() (items []*models.Representative, err error) {
	err = r.db.NewSelect().
		Model(&items).
		Relation("Human").
		Order("human.surname").
		Order("human.name").
		Order("human.patronymic").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Representative, error) {
	item := models.Representative{}
	err := r.db.NewSelect().Model(&item).
		Relation("Human").
		Relation("RepresentativeToPatient.Patient.Human").
		Relation("RepresentativeToPatient.RepresentativeType").
		Where("representative.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Representative{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Representative) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
