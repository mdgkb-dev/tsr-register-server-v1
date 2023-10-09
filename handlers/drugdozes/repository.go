package drugdozes

import (
	"context"
	"fmt"
	"mdgkb/tsr-tegister-server-v1/models"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) Create(c context.Context, item *models.DrugDoze) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll(c context.Context) (item models.DrugDozesWithCount, err error) {
	fmt.Println("r", r)
	item.DrugDozes = make(models.DrugDozes, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&item.DrugDozes)

	r.helper.SQL.ExtractQueryFilter(c).HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) Get(c context.Context, slug string) (*models.DrugDoze, error) {
	item := models.DrugDoze{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("DrugDozeComponents").
		Relation("DrugRegimens.DrugRegimenBlocks.Formula").
		Where("?TableAlias.id = ?", slug).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.DrugDoze{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(c context.Context, item *models.DrugDoze) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
