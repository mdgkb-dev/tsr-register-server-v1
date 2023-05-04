package registers

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

func (r *Repository) create(item *models.Register) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items []*models.Register, err error) {
	err = r.db().NewSelect().
		Model(&items).
		//Relation("RegisterDiagnosis").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.Register, error) {
	item := models.Register{}
	err := r.db().NewSelect().
		Model(&item).
		//Relation("RegisterDiagnosis.MkbDiagnosis.MkbSubDiagnosis").
		//Relation("RegisterDiagnosis.MkbDiagnosis.MkbGroup").
		//Relation("RegisterDiagnosis.MkbSubDiagnosis").
		//Relation("RegisterDiagnosis.MkbConcreteDiagnosis").
		//Relation("Researches", func(q *bun.SelectQuery) *bun.SelectQuery {
		//	return q.Order("register_group.register_group_order")
		//}).
		//Relation("Researches.Questions", func(q *bun.SelectQuery) *bun.SelectQuery {
		//	return q.Order("register_property.register_property_order")
		//}).
		//Relation("Researches.Questions.QuestionExamples").
		//Relation("Researches.Questions.QuestionVariants").
		//Relation("Researches.Questions.ValueType").
		//Relation("Researches.Questions.QuestionMeasures").
		//Relation("Researches.Questions.RegisterPropertySets", func(q *bun.SelectQuery) *bun.SelectQuery {
		//	return q.Order("register_property_set.register_property_set_order")
		//}).
		//Relation("Researches.Questions.RegisterPropertySets.RegisterPropertyOthers").
		//Relation("Researches.Questions.AnswerVariants", func(q *bun.SelectQuery) *bun.SelectQuery {
		//	return q.Order("register_property_radio.register_property_radio_order")
		//}).
		//Relation("Researches.Questions.AnswerVariants.RegisterPropertyOthers", func(q *bun.SelectQuery) *bun.SelectQuery {
		//	return q.Order("register_property_others.register_property_others_order")
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
	//item.RegisterToPatientCount, err = r.db().NewSelect().Model((*models.ResearchResult)(nil)).Where("register_id = ?", id).Count(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Register{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}
func (r *Repository) update(item *models.Register) (err error) {
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
