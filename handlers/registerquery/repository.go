package registerquery

import (
	"fmt"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(query *models.ResearchQuery) (err error) {
	_, err = r.db().NewInsert().Model(query).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (queries models.ResearchQueries, err error) {
	err = r.db().NewSelect().
		Model(&queries).
		Relation("ResearchesPool").
		Scan(r.ctx)
	return queries, err
}

func (r *Repository) get(id string) (*models.ResearchQuery, error) {
	query := models.ResearchQuery{}
	err := r.db().NewSelect().
		Model(&query).
		Relation("ResearchQueryGroups", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("register_query_groups.item_order")
		}).
		Relation("ResearchQueryGroups.ResearchQueryGroupQuestions", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("register_query_group_properties.item_order")
		}).
		Relation("ResearchQueryGroups.ResearchQueryGroupQuestions.Question.ValueType").
		Relation("ResearchQueryGroups.ResearchQueryGroupQuestions.Question.RegisterPropertySets.RegisterPropertyOthers").
		Relation("ResearchQueryGroups.ResearchQueryGroupQuestions.Question.AnswerVariants.RegisterPropertyOthers").
		Relation("ResearchQueryGroups.Researches.RegisterGroupsToPatients.Patient.Human", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("patient.full_name")
		}).
		Relation("ResearchQueryGroups.Researches.RegisterGroupsToPatients.PatientAnswerComments").
		Relation("ResearchQueryGroups.Researches.RegisterGroupsToPatients.Answer.Question.ValueType").
		Relation("ResearchQueryGroups.Researches.RegisterGroupsToPatients.Answer").
		Relation("ResearchesPool.ResearchResult.Patient", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("patient.full_name")
		}).
		Relation("ResearchesPool.ResearchResult.Patient.Human").
		//Relation("ResearchesPool.ResearchResult.Patient.RegisterGroupsToPatient", func(q *bun.SelectQuery) *bun.SelectQuery {
		//	return q.Join("JOIN register_query_groups rqg on rqg.id = register_groups_to_patients.register_group_id").
		//		Order("rqg.item_order")
		//}).
		//Relation("ResearchesPool.ResearchResult.Patient.RegisterGroupsToPatient.Answer", func(q *bun.SelectQuery) *bun.SelectQuery {
		//	return q.Join("JOIN register_query_group_properties rqgp on rqgp.id = register_groups_to_patients.register_property_id").
		//		Order("rqgp.item_order")
		//}).
		Where("register_queries.id = ?", id).Scan(r.ctx)
	return &query, err
}

func (r *Repository) update(query *models.ResearchQuery) (err error) {
	_, err = r.db().NewUpdate().Model(query).Where("id = ?", query.ID).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.ResearchQuery{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) execute(registerQuery *models.ResearchQuery) error {
	fmt.Println(registerQuery)
	return nil
}
