package researchesresults

import (
	"context"
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

func (r *Repository) Create(item *models.ResearchResult) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll() (item models.ResearchResultsWithCount, err error) {
	item.ResearchResults = make(models.ResearchResults, 0)
	query := r.DB().NewSelect().Model(&item.ResearchResults).Relation("Answers")

	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) Get(slug string) (*models.ResearchResult, error) {
	item := models.ResearchResult{}
	err := r.DB().NewSelect().Model(&item).
		Relation("Answers.SelectedAnswerVariants").
		Relation("Answers.AnswerFiles.FileInfo").
		Where("?TableAlias.id = ?", slug).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) Delete(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.ResearchResult{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(item *models.ResearchResult) (err error) {
	_, err = r.DB().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) GetActualAnthropomethryResult(c context.Context, patientID string) (*models.ResearchResult, error) {
	item := models.ResearchResult{}
	query := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Answers.Question").
		Join("join patients_researches pr on  pr.id = ?TableAlias.patient_research_id").
		Join("join researches r on r.id = pr.research_id").
		Join("join questions q on q.research_id = r.id and q.code in (?)", bun.In([]string{string(models.AnthropomethryKeyWeight), string(models.AnthropomethryKeyHeight)})).
		Order("research_results.item_date desc").
		Where("pr.patient_id = ?", patientID).
		Limit(1)
	err := query.Scan(c)
	return &item, err
}
