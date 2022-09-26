package register

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"mdgkb/tsr-tegister-server-v1/models"

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

func (r *Repository) getAll(userID uuid.UUID) (items []*models.Register, err error) {
	err = r.db().NewSelect().
		Model(&items).
		Join("JOIN registers_users ON registers_users.user_id = ? and register.id = registers_users.register_id", userID).
		Relation("RegisterDiagnosis").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.Register, error) {
	item := models.Register{}
	err := r.db().NewSelect().
		Model(&item).
		Relation("RegisterDiagnosis.MkbDiagnosis.MkbSubDiagnosis").
		Relation("RegisterDiagnosis.MkbDiagnosis.MkbGroup").
		Relation("RegisterDiagnosis.MkbSubDiagnosis").
		Relation("RegisterDiagnosis.MkbConcreteDiagnosis").
		Relation("RegisterGroups", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("register_group.register_group_order")
		}).
		Relation("RegisterGroups.RegisterProperties", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("register_property.register_property_order")
		}).
		Relation("RegisterGroups.RegisterProperties.RegisterPropertyExamples").
		Relation("RegisterGroups.RegisterProperties.RegisterPropertyVariants").
		Relation("RegisterGroups.RegisterProperties.ValueType").
		Relation("RegisterGroups.RegisterProperties.RegisterPropertyMeasures").
		Relation("RegisterGroups.RegisterProperties.RegisterPropertySets", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("register_property_set.register_property_set_order")
		}).
		Relation("RegisterGroups.RegisterProperties.RegisterPropertySets.RegisterPropertyOthers").
		Relation("RegisterGroups.RegisterProperties.RegisterPropertyRadios", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("register_property_radio.register_property_radio_order")
		}).
		Relation("RegisterGroups.RegisterProperties.RegisterPropertyRadios.RegisterPropertyOthers", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("register_property_others.register_property_others_order")
		}).
		Relation("RegisterToPatient.Patient.Human", func(q *bun.SelectQuery) *bun.SelectQuery {
			//r.queryFilter.HandleQuery(q)
			return q.Order("patient__human.surname", "patient__human.name", "patient__human.patronymic")
		}).
		//Relation("RegisterToPatient.Patient.RegisterPropertyToPatient.RegisterProperty").
		//Relation("RegisterToPatient.Patient.RegisterPropertyToPatient.RegisterProperty").
		//Relation("RegisterToPatient.Patient.RegisterPropertySetToPatient.RegisterPropertySet").
		Where("register.id = ?", id).Scan(r.ctx)
	if err != nil {
		return nil, err
	}
	item.RegisterToPatientCount, err = r.db().NewSelect().Model((*models.RegisterToPatient)(nil)).Where("register_id = ?", id).Count(r.ctx)
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
