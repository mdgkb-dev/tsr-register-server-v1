package users

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	//_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getAll() (item models.UsersWithCount, err error) {
	item.Users = make(models.Users, 0)
	query := r.db().NewSelect().
		Model(&item.Users).
		Relation("Human").
		Relation("Role")

	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) get(id string) (*models.User, error) {
	item := models.User{}
	err := r.db().NewSelect().
		Model(&item).
		Where("?TableAlias.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) getByUserAccountID(accountId string) (*models.User, error) {
	item := models.User{}
	err := r.db().NewSelect().
		Model(&item).
		Relation("UsersDomains").
		Where("?TableAlias.user_account_id = ?", accountId).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(user *models.User) (err error) {
	_, err = r.db().NewInsert().Model(user).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.User) (err error) {
	_, err = r.db().NewUpdate().Model(item).
		OmitZero().
		ExcludeColumn("password", "is_active"). // all columns except col1
		Where("id = ?", item.ID).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertEmail(item *models.User) (err error) {
	_, err = r.db().NewInsert().On("conflict (email) DO UPDATE").
		Set("phone = EXCLUDED.phone").
		Set("login = ''").
		Model(item).
		Exec(r.ctx)
	return err
}

func (r *Repository) addToUser(values map[string]interface{}, table string) error {
	_, err := r.db().NewInsert().Model(&values).TableExpr(table).Exec(r.ctx)
	return err
}

func (r *Repository) removeFromUser(values map[string]interface{}, table string) error {
	q := r.db().NewDelete().Table(table)
	for key, value := range values {
		q = q.Where("? = ?", bun.Ident(key), value)
	}
	_, err := q.Exec(r.ctx)
	return err
}
