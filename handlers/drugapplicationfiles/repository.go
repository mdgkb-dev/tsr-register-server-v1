package drugapplicationfiles

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"

	"github.com/uptrace/bun"
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

func (r *Repository) Create(item *models.DrugApplicationFile) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll() (items models.DrugApplicationFilesWithCount, err error) {
	items.DrugApplicationFiles = make(models.DrugApplicationFiles, 0)
	query := r.DB().NewSelect().
		Model(&items.DrugApplicationFiles)
	items.Count, err = query.ScanAndCount(r.ctx)
	return items, err
}

func (r *Repository) Get(id string) (*models.DrugApplicationFile, error) {
	item := models.DrugApplicationFile{}
	err := r.DB().NewSelect().
		Model(&item).
		Where("?TableAlias.id = ?", id).Scan(r.ctx)
	if err != nil {
		return nil, err
	}
	return &item, err
}

func (r *Repository) Delete(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.DrugApplicationFile{}).Where("id = ?", id).Exec(r.ctx)
	return err
}
func (r *Repository) Update(item *models.DrugApplicationFile) (err error) {
	_, err = r.DB().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
