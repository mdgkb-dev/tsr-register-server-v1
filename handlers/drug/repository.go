package drug

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Drug) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() ([]*models.Drug, error) {
	items := make([]*models.Drug, 0)
	err := r.db.NewSelect().Model(&items).
		// Relation("DrugRegimens").
		// Relation("DrugRegimens.DrugRegimenBlocks").
		// Relation("DrugRegimens.DrugRegimenBlocks.DrugRegimenBlockItems").
		// Order("DrugRegimens.DrugRegimenBlocks.OrderItem").
		// Order("DrugRegimens.DrugRegimenBlockItems.OrderItem").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Drug, error) {
	item := models.Drug{}
	err := r.db.NewSelect().Model(&item).
		// Relation("DrugRegimens.DrugRegimenBlocks.DrugRegimenBlockItems").
		Relation("DrugRegimens.DrugRegimenBlocks.DrugRegimenBlockItems", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("drug_regimen_block_items.order_item")
		}).
		Relation("DrugRegimens.DrugRegimenBlocks", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("drug_regimen_blocks.order_item")
		}).
// Order("drug_regimens.drug_regimen_blocks.order_item").
		// Order("DrugRegimens.DrugRegimenBlockItems.OrderItem").
		Where("id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Drug{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Drug) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
