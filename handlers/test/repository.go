package test

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

type IRepository interface {
	create(*gin.Context, *models.Test) error
	getAll(*gin.Context) ([]models.Test, error)
	get(*gin.Context, string) (models.Test, error)
	updateStatus(*gin.Context, *models.Test) error
	delete(*gin.Context, string) error
	update(*gin.Context, *models.Test) error
}

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) create(ctx *gin.Context, item *models.Test) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(ctx)
	return err
}

func (r *Repository) getAll(ctx *gin.Context) (items []models.Test, err error) {
	err = r.db.NewSelect().Model(&items).Scan(ctx)
	return items, err
}

func (r *Repository) get(ctx *gin.Context, id string) (item models.Test, err error) {
	err = r.db.NewSelect().
		Model(&item).
		Relation("Timetable.TimetableDays.Weekday").
		Relation("Schedule.ScheduleItems").
		Where("Test.id = ?", id).Scan(ctx)
	return item, err
}

func (r *Repository) updateStatus(ctx *gin.Context, item *models.Test) (err error) {
	_, err = r.db.NewUpdate().Model(item).Exec(ctx)
	return err
}

func (r *Repository) delete(ctx *gin.Context, id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Test{}).Where("id = ?", id).Exec(ctx)
	return err
}

func (r *Repository) update(ctx *gin.Context, item *models.Test) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(ctx)
	return err
}
