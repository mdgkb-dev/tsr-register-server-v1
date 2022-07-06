package drugRegimenBlockItem

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items []*models.DrugRegimenBlockItem) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.DrugRegimenBlockItem)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items []*models.DrugRegimenBlockItem) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("drug_regimen_block_id = EXCLUDED.drug_regimen_block_id").
		Set("days_count = EXCLUDED.days_count").
		Set("times_per_day = EXCLUDED.times_per_day").
		Set("order_item = EXCLUDED.order_item").
		Model(&items).
		Exec(r.ctx)
	return err
}
