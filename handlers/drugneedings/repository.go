package drugneedings

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) Create(c context.Context, item *models.DrugNeeding) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.DrugNeeding{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(c context.Context, item *models.DrugNeeding) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) Upsert(c context.Context, item *models.DrugNeeding) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) Get(c context.Context, id string) (*models.DrugNeeding, error) {
	item := models.DrugNeeding{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("DrugNeedingBlocks").
		Where("?TableAlias.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) GetAll(c context.Context) (item models.DrugNeedingsWithCount, err error) {
	item.DrugNeedings = make(models.DrugNeedings, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&item.DrugNeedings)
	r.helper.SQL.ExtractQueryFilter(c).HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) CreateMany(c context.Context, items models.DrugNeedings) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) DeleteMany(c context.Context, idPool []uuid.UUID) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().
		Model((*models.DrugNeeding)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) UpsertMany(c context.Context, items models.DrugNeedings) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").
		Set("drug_id = EXCLUDED.drug_id").
		Set("name = EXCLUDED.name").
		Model(&items).
		Exec(r.ctx)
	return err
}
