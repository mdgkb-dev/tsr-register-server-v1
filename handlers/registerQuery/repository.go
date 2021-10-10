package registerQuery

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(query *models.RegisterQuery) (err error) {
	_, err = r.db.NewInsert().Model(query).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (queries models.RegisterQueries, err error) {
	err = r.db.NewSelect().Model(&queries).Scan(r.ctx)
	return queries, err
}

func (r *Repository) get(id *string) (*models.RegisterQuery, error) {
	query := models.RegisterQuery{}
	err := r.db.NewSelect().
		Model(&query).
		Relation("RegisterQueryToRegisterProperty.RegisterProperty").
		Where("register_queries.id = ?", *id).Scan(r.ctx)
	return &query, err
}

func (r *Repository) update(query *models.RegisterQuery) (err error) {
	_, err = r.db.NewUpdate().Model(query).Where("id = ?", query.ID).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.RegisterQuery{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}
