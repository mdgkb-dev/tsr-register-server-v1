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
	err = r.db.NewSelect().
		Model(&queries).
		Relation("Register").
		Scan(r.ctx)
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

func (r *Repository) execute(registerQuery *models.RegisterQuery) ([]map[string]interface{}, error) {
	result := make([]map[string]interface{}, 0)

	query := `SELECT *
		FROM   crosstab(
	$$
	select
		rptp.patient_id, rp.name, rptp.value_string
		from
		register_queries
		join register_query_to_register_property rqtrp on register_queries.id = rqtrp.register_query_id
		join register_property rp on rqtrp.register_property_id = rp.id
		join register_property_to_patient rptp on rp.id = rptp.register_property_id
		order by 1, 2, 3
		;
		$$
	) AS ct ("in" uuid, "1" varchar, "2" varchar, "3" varchar, "4" varchar, "5" varchar);
`

	res, err := r.db.QueryContext(r.ctx, query, result)
	err = r.db.ScanRows(r.ctx, res, &result)
	return result, err
}
