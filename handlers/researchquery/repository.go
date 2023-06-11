package researchquery

import (
	"fmt"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func (r *Repository) DB() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) SetQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Create(query *models.ResearchQuery) (err error) {
	_, err = r.DB().NewInsert().Model(query).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll() (queries models.ResearchQueriesWithCount, err error) {
	err = r.DB().NewSelect().
		Model(&queries).
		Relation("ResearchesPool").
		Scan(r.ctx)
	return queries, err
}

func (r *Repository) Get(id string) (*models.ResearchQuery, error) {
	query := models.ResearchQuery{}
	err := r.DB().NewSelect().
		Model(&query).
		Relation("ResearchQueryGroups", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("research_query_groups.item_order")
		}).
		Relation("ResearchQueryGroups.ResearchQueryGroupQuestions", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("research_query_group_questions.item_order")
		}).
		Relation("ResearchQueryGroups.ResearchQueryGroupQuestions.Question.ValueType").
		Relation("ResearchQueryGroups.ResearchQueryGroupQuestions.Question.AnswerVariants", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("answer_variants.item_order")
		}).
		Relation("ResearchQueryGroups.ResearchQueryGroupQuestions.Question.Children.ValueType").
		//Relation("ResearchQueryGroups.ResearchQueryGroupQuestions.Question.AnswerVariants.RegisterPropertyOthers").
		//Relation("ResearchQueryGroups.Researches.RegisterGroupsToPatients.Patient.Human", func(q *bun.SelectQuery) *bun.SelectQuery {
		//	return q.Order("patients.full_name")
		//}).
		Relation("ResearchQueryGroups.Research.ResearchResults", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Join("JOIN patients_view p on research_results.patient_id = p.id").
				Order("p.full_name")
		}).
		Relation("ResearchQueryGroups.Research.ResearchResults.PatientResearch").
		Relation("ResearchQueryGroups.Research.ResearchResults.PatientResearch").
		Relation("ResearchQueryGroups.Research.ResearchResults.Answers.Question.ValueType").
		Relation("ResearchQueryGroups.Research.ResearchResults.Answers.SelectedAnswerVariants").
		//Relation("ResearchQueryGroups.Researches.RegisterGroupsToPatients.PatientAnswerComments").
		//Relation("ResearchQueryGroups.Researches.RegisterGroupsToPatients.Answer.Question.ValueType").
		//Relation("ResearchQueryGroups.Researches.RegisterGroupsToPatients.Answer").
		Relation("ResearchesPool.PatientsResearchesPools.Patient", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("patient.full_name")
		}).
		Relation("ResearchesPool.PatientsResearchesPools.Patient.Human").
		//Relation("ResearchesPool.PatientsResearchesPools.Patient.PatientsResearches", func(q *bun.SelectQuery) *bun.SelectQuery {
		//	return q.Join("JOIN research_query_groups rqg on rqg.id = patients_resear.research_group_id").
		//		Order("rqg.item_order")
		//}).
		Relation("ResearchesPool.PatientsResearchesPools.Patient.PatientsResearches").
		//Relation("ResearchesPool.ResearchResult.Patient.RegisterGroupsToPatient.Answer", func(q *bun.SelectQuery) *bun.SelectQuery {
		//	return q.Join("JOIN research_query_group_properties rqgp on rqgp.id = research_groups_to_patients.research_property_id").
		//		Order("rqgp.item_order")
		//}).
		Where("?TableAlias.id = ?", id).Scan(r.ctx)
	return &query, err
}

func (r *Repository) Update(query *models.ResearchQuery) (err error) {
	_, err = r.DB().NewUpdate().Model(query).Where("id = ?", query.ID).Exec(r.ctx)
	return err
}

func (r *Repository) Delete(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.ResearchQuery{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Execute(registerQuery *models.ResearchQuery) error {
	fmt.Println(registerQuery)
	return nil
}
