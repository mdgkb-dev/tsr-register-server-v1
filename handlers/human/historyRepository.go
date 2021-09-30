package human

import (
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *HistoryRepository) getDB() *bun.DB {
	return r.db
}

func (r *HistoryRepository) create(item *models.HumanHistory) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}
