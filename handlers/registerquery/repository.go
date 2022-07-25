package registerquery

import (
	"fmt"
	"mdgkb/tsr-tegister-server-v1/models"
	"strings"

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
		Relation("RegisterQueryToRegisterProperty.RegisterProperty.ValueType").
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
	cols := make([]string, 0)
	values := make([]string, 0)
	for i, regQueryToRegProp := range registerQuery.RegisterQueryToRegisterProperty {
		valueType := "varchar"
		if regQueryToRegProp.RegisterProperty.ValueType.Name == "number" {
			valueType = "int"
		}
		if regQueryToRegProp.RegisterProperty.ValueType.Name == "date" {
			valueType = "date"
		}
		name := "\"" + regQueryToRegProp.RegisterProperty.ShortName + "\" "
		name1 := "'" + regQueryToRegProp.RegisterProperty.ShortName + "' "
		nameWithType := name + valueType
		nameForValues := "(" + name1 + ")"

		if i < len(registerQuery.RegisterQueryToRegisterProperty)-1 {
			nameWithType = nameWithType + ","
			nameForValues = nameForValues + ","
		}
		cols = append(cols, nameWithType)
		values = append(values, nameForValues)
	}

	colsString := strings.Join(cols, "  ")
	valuesString := strings.Join(values, "  ")

	query := fmt.Sprintf(`SELECT *
		FROM   public.crosstab(
	    $$
  select
                     case
             when vt.name::varchar = 'set' then concat_ws(' '::text, h1.surname, h1.name, h1.patronymic)
             else
              concat_ws(' '::text, h.surname, h.name, h.patronymic)
                 end,
              rp.short_name,
  CASE
  WHEN vt.name::varchar = 'string' then rptp.value_string
  WHEN vt.name::varchar = 'text' then rptp.value_string
  WHEN vt.name::varchar = 'number' then rptp.value_number::varchar
  WHEN vt.name::varchar = 'date' then rptp.value_date::varchar
  WHEN vt.name::varchar = 'radio' then rpr.name
  WHEN vt.name::varchar = 'set' then rps.name
  END
            from
            register_queries
            join register_query_to_register_property rqtrp on register_queries.id = rqtrp.register_query_id
            join register_property rp on rqtrp.register_property_id = rp.id
            join value_type vt on vt.id = rp.value_type_id 
            left join register_property_to_patient rptp on rp.id = rptp.register_property_id
            left join patients p on p.id = rptp.patient_id
            left join human h on h.id = p.human_id
            left join register_property_radio rpr on rpr.id = rptp.register_property_radio_id
            left join register_property_set rps on rps.register_property_id = rp.id
            left join register_property_set_to_patient rpstp on rpstp.register_property_set_id = rps.id

             left join patients p1 on p1.id = rpstp.patient_id
             left join human h1 on h1.id = p1.human_id
            order by 1, 2
            ;
		$$, $$values %s $$
	) AS ct ("%s" varchar, %s);
`, valuesString, registerQuery.Key, colsString)

	res, err := r.db().QueryContext(r.ctx, query, &registerQuery.Data)
	if err != nil {
		return err
	}
	err = r.db().ScanRows(r.ctx, res, &registerQuery.Data)

	return err
}
