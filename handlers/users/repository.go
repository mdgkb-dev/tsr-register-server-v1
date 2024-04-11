package users

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) GetAll(c context.Context) (item models.UsersWithCount, err error) {
	item.Users = make(models.Users, 0)
	q := r.helper.DB.IDB(c).NewSelect().
		Model(&item.Users).
		Relation("UserAccount")

	r.helper.SQL.ExtractFTSP(c).HandleQuery(q)
	item.Count, err = q.ScanAndCount(c)
	return item, err
}

func (r *Repository) Get(c context.Context, id string) (*models.User, error) {
	item := models.User{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Where("", id).
		Scan(c)
	return &item, err
}

func (r *Repository) GetByUserAccountID(c context.Context, id string) (*models.User, error) {
	item := models.User{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("UsersDomains").
		Where("?TableAlias.user_account_id = ?", id).
		Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.User{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.User) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) Create(c context.Context, item *models.User) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}
