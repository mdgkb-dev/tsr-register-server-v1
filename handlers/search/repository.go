package search

import (
	"context"
	"fmt"
	"mdgkb/tsr-tegister-server-v1/middleware"

	"github.com/pro-assistance/pro-assister/search"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) getGroups(groupID string) (search.SearchGroups, error) {
	items := make(search.SearchGroups, 0)
	query := r.db().NewSelect().Model(&items).
		Relation("SearchGroupMetaColumns").
		Order("search_group_order")

	if groupID != "" {
		query = query.Where("id = ?", groupID)
	}
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) search(c context.Context, searchModel *search.SearchModel) error {
	querySelect := fmt.Sprintf("SELECT %s.%s as value, substring(%s for 40) as label", searchModel.SearchGroup.Table, searchModel.SearchGroup.ValueColumn, searchModel.SearchGroup.LabelColumn)
	queryFrom := fmt.Sprintf("FROM %s", searchModel.SearchGroup.Table)

	join := fmt.Sprintf("JOIN %s on %s.%s = %s.id and %s.domain_id in (?)", searchModel.SearchGroup.JoinTable, searchModel.SearchGroup.JoinTable, searchModel.SearchGroup.JoinColumn, searchModel.SearchGroup.Table, searchModel.SearchGroup.JoinTable)

	condition := fmt.Sprintf("where regexp_replace(%s, '[^а-яА-Яa-zA-Z0-9. ]', '', 'g') ILIKE %s", searchModel.SearchGroup.SearchColumn, "'%"+searchModel.Query+"%'")
	conditionTranslitToRu := fmt.Sprintf("or regexp_replace(%s, '[^а-яА-Яa-zA-Z0-9. ]', '', 'g') ILIKE %s", searchModel.SearchGroup.SearchColumn, "'%"+r.helper.Util.TranslitToRu(searchModel.Query)+"%'")
	conditionTranslitToEng := fmt.Sprintf("or regexp_replace(%s, '[^а-яА-Яa-zA-Z0-9. ]', '', 'g') ILIKE %s", searchModel.SearchGroup.SearchColumn, "'%"+r.helper.Util.TranslitToEng(searchModel.Query)+"%'")

	queryOrder := fmt.Sprintf("ORDER BY %s", searchModel.SearchGroup.LabelColumn)
	query := fmt.Sprintf("%s %s %s %s %s %s %s", querySelect, queryFrom, join, condition, conditionTranslitToRu, conditionTranslitToEng, queryOrder)
	rows, err := r.db().QueryContext(r.ctx, query, bun.In(middleware.ClaimDomainIDS.FromContextSlice(c)))
	if err != nil {
		return err
	}

	err = r.db().ScanRows(r.ctx, rows, &searchModel.SearchGroup.SearchElements)
	return err
}

func (r *Repository) elasticSearch(model *search.SearchModel) error {
	var data map[string]interface{}
	//query, indexes := model.BuildQuery()
	//res, err := r.elasticsearch.Search(
	//	r.elasticsearch.Search.WithIndex(indexes...),
	//	r.elasticsearch.Search.WithBody(esutil.NewJSONReader(&query)),
	//	r.elasticsearch.Search.WithPretty(),
	//)
	//defer res.Body.Close()
	//if err != nil {
	//	return err
	//}
	//err = json.NewDecoder(res.Body).Decode(&data)
	//if err != nil {
	//	return err
	//}
	model.ParseMap(data)
	return nil
}

func (r *Repository) elasticSuggester(model *search.SearchModel) error {
	var re map[string]interface{}
	//indexes := []string{}
	//if model.SearchGroup != nil {
	//	indexes = append(indexes, model.SearchGroup.Table)
	//}
	//should := make([]interface{}, 0)
	//should = append(should, map[string]interface{}{
	//	"prefix": map[string]interface{}{
	//		"name": map[string]interface{}{
	//			"value": model.Query,
	//		},
	//	},
	//})
	//var buf bytes.Buffer
	//query := map[string]interface{}{
	//	"query": should[0],
	//}
	//_ = json.NewEncoder(&buf).Encode(query)
	//res, err := r.elasticsearch.Search(
	//	//r.elasticsearch.Search.WithIndex(indexes...),
	//	r.elasticsearch.Search.WithBody(&buf),
	//	//r.elasticsearch.Get.
	//	r.elasticsearch.Search.WithPretty(),
	//)
	//defer res.Body.Close()
	//if err != nil {
	//	return err
	//}
	//err = json.NewDecoder(res.Body).Decode(&re)
	//if err != nil {
	//	return err
	//}
	model.ParseMap(re)
	return nil
}
