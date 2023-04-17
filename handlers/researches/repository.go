package researches

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"

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

func (r *Repository) create(item *models.ResearchesPool) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items models.ResearchesPools, err error) {
	err = r.db().NewSelect().
		Model(&items).
		Relation("ResearchDiagnosis").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.ResearchesPool, error) {
	item := models.ResearchesPool{}
	err := r.db().NewSelect().
		Model(&item).
		//Relation("ResearchDiagnosis.MkbDiagnosis.MkbSubDiagnosis").
		//Relation("ResearchDiagnosis.MkbDiagnosis.MkbGroup").
		//Relation("ResearchDiagnosis.MkbSubDiagnosis").
		//Relation("ResearchDiagnosis.MkbConcreteDiagnosis").
		Relation("ResearchSections", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("research_sections.item_order")
		}).
		Relation("ResearchSections.Questions", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("research_questions.item_order")
		}).
		Relation("ResearchSections.Questions.QuestionExamples").
		Relation("ResearchSections.Questions.QuestionVariants").
		Relation("ResearchSections.Questions.ValueType").
		Relation("ResearchSections.Questions.QuestionMeasures").
		//Relation("ResearchSections.Questions.ResearchPropertySets", func(q *bun.SelectQuery) *bun.SelectQuery {
		//	return q.Order("Research_property_set.Research_property_set_order")
		//}).
		//Relation("ResearchSections.Questions.ResearchPropertySets.ResearchPropertyOthers").
		//Relation("ResearchSections.Questions.AnswersVariants", func(q *bun.SelectQuery) *bun.SelectQuery {
		//	return q.Order("Research_property_radio.Research_property_radio_order")
		//}).
		//Relation("ResearchSections.Questions.AnswersVariants.ResearchPropertyOthers", func(q *bun.SelectQuery) *bun.SelectQuery {
		//	return q.Order("Research_property_others.Research_property_others_order")
		//}).
		//Relation("ResearchResult.Patient.Human", func(q *bun.SelectQuery) *bun.SelectQuery {
		//	//r.queryFilter.HandleQuery(q)
		//	return q.Order("patient__human.surname", "patient__human.name", "patient__human.patronymic")
		//}).
		//Relation("ResearchResult.Patient.Answer.Question").
		//Relation("ResearchResult.Patient.Answer.Question").
		//Relation("ResearchResult.Patient.Answer.AnswerVariant").
		Where("?TableAlias.id = ?", id).Scan(r.ctx)
	if err != nil {
		return nil, err
	}
	//item.ResearchToPatientCount, err = r.db().NewSelect().Model((*models.ResearchResult)(nil)).Where("Research_id = ?", id).Count(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.ResearchesPool{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}
func (r *Repository) update(item *models.ResearchesPool) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) getValueTypes() (models.ValueTypes, error) {
	items := make(models.ValueTypes, 0)
	err := r.db().NewSelect().
		Model(&items).
		Scan(r.ctx)
	return items, err
}
