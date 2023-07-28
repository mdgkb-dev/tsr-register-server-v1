package documentfieldvalues

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) DB() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) SetQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Create(item *models.DocumentFieldValue) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll() (item models.DocumentFieldValuesWithCount, err error) {
	item.DocumentFieldValues = make(models.DocumentFieldValues, 0)
	query := r.DB().NewSelect().Model(&item.DocumentFieldValues)

	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) Get(slug string) (*models.DocumentFieldValue, error) {
	item := models.DocumentFieldValue{}
	err := r.DB().NewSelect().Model(&item).
		Relation("DocumentFieldValueFieldValues.DocumentFieldValueTypeField.ValueType").
		Where("?TableAlias.id = ?", slug).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) Delete(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.DocumentFieldValue{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(item *models.DocumentFieldValue) (err error) {
	_, err = r.DB().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) UpsertMany(items models.DocumentFieldValues) (err error) {
	_, err = r.DB().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("value_string = EXCLUDED.value_string").
		Set("value_number = EXCLUDED.value_number").
		Set("value_date = EXCLUDED.value_date").
		Set("document_id = EXCLUDED.document_id").
		Exec(r.ctx)
	return err
}
