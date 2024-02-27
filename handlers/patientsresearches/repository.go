package patientsresearches

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

func (r *Repository) Create(item *models.PatientResearch) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll() (item models.PatientsResearchesWithCount, err error) {
	item.PatientsResearches = make(models.PatientsResearches, 0)
	query := r.DB().NewSelect().Model(&item.PatientsResearches)

	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) Get(slug string) (*models.PatientResearch, error) {
	item := models.PatientResearch{}
	err := r.DB().NewSelect().Model(&item).
		Where("?TableAlias.id = ?", slug).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) Delete(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.PatientResearch{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(item *models.PatientResearch) (err error) {
	_, err = r.DB().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) GetPatientResearch(c context.Context, patientId string, researchId string) (*models.PatientResearch, error) {
	item := models.PatientResearch{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("ResearchResults.Answers.SelectedAnswerVariants").
		Relation("ResearchResults.Answers.AnswerFiles.FileInfo").
		Where("?TableAlias.patient_id  = ?", patientId).
		Where("?TableAlias.research_id = ?", researchId).
		Scan(c)
	return &item, err
}
