package questionvariants

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) GetAll(c context.Context) (item models.QuestionVariantsWithCount, err error) {
	item.QuestionVariants = make(models.QuestionVariants, 0)
	q := r.helper.DB.IDB(c).NewSelect().Model(&item.QuestionVariants)

	r.helper.SQL.ExtractFTSP(c).HandleQuery(q)
	item.Count, err = q.ScanAndCount(c)
	return item, err
}

func (r *Repository) Get(c context.Context, id string) (*models.QuestionVariant, error) {
	item := models.QuestionVariant{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Where("?TableAlias.id = ?", id).
		Scan(c)
	return &item, err
}

func (r *Repository) GetByQuestionVariantAccountID(c context.Context, id string) (*models.QuestionVariant, error) {
	item := models.QuestionVariant{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("QuestionVariantsDomains").
		Where("?TableAlias.user_account_id = ?", id).
		Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.QuestionVariant{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.QuestionVariant) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) Create(c context.Context, item *models.QuestionVariant) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}
