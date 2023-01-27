package registerquery

import (
	"fmt"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(query *models.RegisterQuery) (err error) {
	_, err = r.db().NewInsert().Model(query).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (queries models.RegisterQueries, err error) {
	err = r.db().NewSelect().
		Model(&queries).
		Relation("Register").
		Scan(r.ctx)
	return queries, err
}

func (r *Repository) get(id string) (*models.RegisterQuery, error) {
	query := models.RegisterQuery{}
	err := r.db().NewSelect().
		Model(&query).
		Relation("RegisterQueryGroups", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("register_query_groups.item_order")
		}).
		Relation("RegisterQueryGroups.RegisterQueryGroupProperties", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("register_query_group_properties.item_order")
		}).
		Relation("RegisterQueryGroups.RegisterQueryGroupProperties.RegisterProperty.ValueType").
		Relation("RegisterQueryGroups.RegisterQueryGroupProperties.RegisterProperty.RegisterPropertySets.RegisterPropertyOthers").
		Relation("RegisterQueryGroups.RegisterQueryGroupProperties.RegisterProperty.RegisterPropertyRadios.RegisterPropertyOthers").
		Relation("RegisterQueryGroups.RegisterGroup.RegisterGroupsToPatients.Patient.Human", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("patient.full_name")
		}).
		Relation("RegisterQueryGroups.RegisterGroup.RegisterGroupsToPatients.RegisterPropertyOthersToPatient").
		Relation("RegisterQueryGroups.RegisterGroup.RegisterGroupsToPatients.RegisterPropertyToPatient.RegisterProperty.ValueType").
		Relation("RegisterQueryGroups.RegisterGroup.RegisterGroupsToPatients.RegisterPropertySetToPatient").
		Relation("Register.RegisterToPatient.Patient", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("patient.full_name")
		}).
		Relation("Register.RegisterToPatient.Patient.Human").
		//Relation("Register.RegisterToPatient.Patient.RegisterGroupsToPatient", func(q *bun.SelectQuery) *bun.SelectQuery {
		//	return q.Join("JOIN register_query_groups rqg on rqg.id = register_groups_to_patients.register_group_id").
		//		Order("rqg.item_order")
		//}).
		//Relation("Register.RegisterToPatient.Patient.RegisterGroupsToPatient.RegisterPropertyToPatient", func(q *bun.SelectQuery) *bun.SelectQuery {
		//	return q.Join("JOIN register_query_group_properties rqgp on rqgp.id = register_groups_to_patients.register_property_id").
		//		Order("rqgp.item_order")
		//}).
		Where("register_queries.id = ?", id).Scan(r.ctx)
	return &query, err
}

func (r *Repository) update(query *models.RegisterQuery) (err error) {
	_, err = r.db().NewUpdate().Model(query).Where("id = ?", query.ID).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.RegisterQuery{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) execute(registerQuery *models.RegisterQuery) error {
	fmt.Println(registerQuery)
	return nil
}
