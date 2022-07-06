package human

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *HistoryRepository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *HistoryRepository) create(item *models.HumanHistory) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}
