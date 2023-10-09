package drugregimens

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) Create(c context.Context, item *models.DrugRegimen) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.DrugRegimen{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(c context.Context, item *models.DrugRegimen) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) Upsert(c context.Context, item *models.DrugRegimen) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) Get(c context.Context, id string) (*models.DrugRegimen, error) {
	item := models.DrugRegimen{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("DrugRegimenBlocks").
		Where("?TableAlias.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) GetAll(c context.Context) (item models.DrugRegimensWithCount, err error) {
	item.DrugRegimens = make(models.DrugRegimens, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&item.DrugRegimens)
	r.helper.SQL.ExtractQueryFilter(c).HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) CreateMany(c context.Context, items models.DrugRegimens) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) DeleteMany(c context.Context, idPool []uuid.UUID) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().
		Model((*models.DrugRegimen)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) UpsertMany(c context.Context, items models.DrugRegimens) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").
		Set("drug_id = EXCLUDED.drug_id").
		Set("name = EXCLUDED.name").
		Model(&items).
		Exec(r.ctx)
	return err
}
