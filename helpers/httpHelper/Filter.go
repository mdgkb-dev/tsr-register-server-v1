package httpHelper

import (
	"fmt"
	"github.com/uptrace/bun"
)

//
//// CreateOrder func
//func CreateOrder(tbl *bun.SelectQuery, args ...string) *bun.SelectQuery {
//	var sortModel SortModel
//	err := json.Unmarshal([]byte(args[0]), &sortModel)
//	if err != nil {
//		return tbl
//	}
//	prefix := ""
//	if args[1] != "" {
//		prefix = fmt.Sprintf("%s.", args[1])
//	}
//	for _, sort := range sortModel {
//		if strings.Contains(sort["colId"], ".") {
//			tbl = tbl.Order(fmt.Sprintf("%s %s", sort["colId"], sort["sort"]))
//		} else {
//			tbl = tbl.Order(fmt.Sprintf("%s%s %s", prefix, sort["colId"], sort["sort"]))
//		}
//	}
//	return tbl
//}
//

// CreateFilter func
func CreateFilter(query *bun.SelectQuery, filterModels FilterModels) {
	if len(filterModels) == 0 {
		return
	}
	for _, filter := range filterModels {
		fmt.Println(filter)
		switch *filter.Type {
		//case "set":
		//	if len(*filter.Values) > 0 {
		//		// проверка на наличие в set null значения
		//		var nullExistInSet = false
		//		for _, val := range *filter.Values {
		//			if val == "" {
		//				nullExistInSet = true
		//				break
		//			}
		//		}
		//		if nullExistInSet { // если null есть - то добавляем условие OR _ IS NULL
		//			tbl = tbl.Where(fmt.Sprintf("%s in (?) OR %s IS NULL", field, field), *filter.Values)
		//		} else {
		//			tbl = tbl.Where(fmt.Sprintf("%s in (?)", field), *filter.Values)
		//		}
		//	}
		case DateType:
			filter.DatesToString()
			constructWhere(query, filter)
		case StringType:
			if filter.IsLike() {
				filter.LikeToString()
			}
			constructWhere(query, filter)
		//case "number":
		//	tbl = constructNumberWhere(tbl, field, filter)
		//case "text":
		//	if filterOperator == "" {
		//		tbl = constructTextWhere(tbl, field, filterOperator, filter)
		//	} else {
		//		tbl = constructTextWhere(tbl, field, filterOperator, filter.Condition1.Filter, filter.Condition2.Filter)
		//	}
		default:
			//log.Println("unknown number filterType: " + *filter.FilterType)
			return
		}
	}
	return
}

func constructWhere(query *bun.SelectQuery, filter *FilterModel) {
	q := ""
	if filter.IsUnary() {
		q = fmt.Sprintf("%s %s '%s'", filter.GetTableAndCol(), *filter.Operator, filter.Value1)
	}
	if filter.IsBetween() {
		q = fmt.Sprintf("%s %s '%s' and '%s'", filter.GetTableAndCol(), *filter.Operator, filter.Value1, filter.Value2)
	}
	query = query.Where(q)
}

//
//func constructTextWhere(tbl *bun.SelectQuery, field string, operator string, filters ...models.Filter) *bun.SelectQuery {
//	operators := map[string]string{
//		"equals":      "%s = ?",
//		"notEqual":    "%s <> ?",
//		"contains":    "%s LIKE ?",
//		"notContains": "%s NOT LIKE ?",
//		"startsWith":  "%s LIKE ?",
//		"endsWith":    "%s LIKE ?",
//	}
//	if operator == "" {
//		tbl = constructQuery(tbl, operators[*filters[0].Type], "", field, operator, likeMix(*filters[0].Type, fmt.Sprintf("%v", *filters[0].Filter)))
//	} else {
//		tbl = constructQuery(tbl, operators[*filters[0].Type], operators[*filters[1].Type], field, operator, likeMix(*filters[0].Type, (*filters[0].Filter).(string)), likeMix(*filters[1].Type, (*filters[1].Filter).(string)))
//	}
//	return tbl
//}
//
//func likeMix(typeOperator, filter string) (result string) {
//	likeOperators := map[string]string{
//		"contains":    "%%%s%%",
//		"notContains": "%%%s%%",
//		"startsWith":  "%s%%",
//		"endsWith":    "%%%s",
//	}
//	likePhrase, ok := likeOperators[typeOperator]
//	if ok {
//		return fmt.Sprintf(likePhrase, filter)
//	}
//	return filter
//}
//
//func getTimeString(date string) string {
//	resultDate := time.Now()
//	resultDate, err := time.Parse("2006-01-02 15:04:05", date)
//	if err != nil {
//		resultDate = time.Date(resultDate.Year(), resultDate.Month(), resultDate.Day(), 0, 0, 0, 0, time.UTC)
//	}
//	location, _ := time.LoadLocation("Europe/Moscow")
//	result := resultDate.In(location).Format("2006-01-02")
//	return result
//}
//
