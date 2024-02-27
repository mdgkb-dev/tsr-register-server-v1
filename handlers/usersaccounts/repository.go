package usersaccounts

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

func (r *Repository) Create(item *models.UserAccount) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetByEmail(email string) (*models.UserAccount, error) {
	item := models.UserAccount{}
	err := r.DB().NewSelect().Model(&item).
		Where("?TableAlias.email = ?", email).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) Get(id string) (*models.UserAccount, error) {
	item := models.UserAccount{}
	err := r.DB().NewSelect().Model(&item).
		Where("?TableAlias.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) UpdateUUID(id string) (err error) {
	_, err = r.DB().NewUpdate().
		Model(models.UserAccount{}).
		Set("uuid = uuid_generate_v4()").
		Where("?TableAlias.id = ?", id).
		Exec(r.ctx)
	return err
}

func (r *Repository) UpdatePassword(item *models.UserAccount) (err error) {
	_, err = r.DB().NewUpdate().
		Model(item).
		Set("password = ?", item.Password).
		Where("id = ?", item.ID).
		Exec(r.ctx)
	return err
}
