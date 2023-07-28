package documentfileinfos

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

func (r *Repository) Create(item *models.DocumentFileInfo) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll() (item models.DocumentFileInfosWithCount, err error) {
	item.DocumentFileInfos = make(models.DocumentFileInfos, 0)
	query := r.DB().NewSelect().Model(&item.DocumentFileInfos)

	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) Get(slug string) (*models.DocumentFileInfo, error) {
	item := models.DocumentFileInfo{}
	err := r.DB().NewSelect().Model(&item).
		Relation("DocumentFileInfoFieldValues.DocumentFileInfoTypeField.ValueType").
		Where("?TableAlias.id = ?", slug).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) Delete(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.DocumentFileInfo{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(item *models.DocumentFileInfo) (err error) {
	_, err = r.DB().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) UpsertMany(items models.DocumentFileInfos) (err error) {
	_, err = r.DB().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("file_info_id = EXCLUDED.file_info_id").
		Set("document_id = EXCLUDED.document_id").
		Exec(r.ctx)
	return err
}
