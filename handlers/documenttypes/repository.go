package documenttypes

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
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Create(item *models.DocumentType) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll() (item models.DocumentTypesWithCount, err error) {
	item.DocumentTypes = make(models.DocumentTypes, 0)
	query := r.DB().NewSelect().Model(&item.DocumentTypes).Relation("DocumentTypeFields.ValueType")

	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) Get(slug string) (*models.DocumentType, error) {
	item := models.DocumentType{}
	err := r.DB().NewSelect().Model(&item).
		Relation("DocumentTypeFields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("document_type_fields.item_order")
		}).
		Relation("DocumentTypeFields.ValueType").
		Where("?TableAlias.id = ?", slug).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) Delete(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.DocumentType{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(item *models.DocumentType) (err error) {
	_, err = r.DB().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
