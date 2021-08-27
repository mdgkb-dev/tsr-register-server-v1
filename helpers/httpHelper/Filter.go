package httpHelper

import (
	"encoding/json"
	"fmt"
	"log"
	"mdgkb/tsr-tegister-server-v1/models"
	"strings"
	"time"

	"github.com/uptrace/bun"
)

// CreateOrder func
func CreateOrder(tbl *bun.SelectQuery, args ...string) *bun.SelectQuery {
	var sortModel models.SortModel
	err := json.Unmarshal([]byte(args[0]), &sortModel)
	if err != nil {
		return tbl
	}
	prefix := ""
	if args[1] != "" {
		prefix = fmt.Sprintf("%s.", args[1])
	}
	for _, sort := range sortModel {
		if strings.Contains(sort["colId"], ".") {
			tbl = tbl.Order(fmt.Sprintf("%s %s", sort["colId"], sort["sort"]))
		} else {
			tbl = tbl.Order(fmt.Sprintf("%s%s %s", prefix, sort["colId"], sort["sort"]))
		}
	}
	return tbl
}

// CreateFilter func
func CreateFilter(tbl *bun.SelectQuery, args ...string) *bun.SelectQuery {
	filterModel, err := models.ParseJSONToFilterModel(args[0])
	if err != nil {
		return tbl
	}
	prefix := ""
	if args[1] != "" {
		prefix = fmt.Sprintf("%s.", args[1])
	}
	for field, filter := range filterModel {
		if filter != (models.Filter{}) {
			if !strings.Contains(field, ".") {
				field = fmt.Sprintf("%s%s", prefix, field)
			}
			var filterOperator string
			if filter.Operator != nil {
				filterOperator = *filter.Operator
			}

			switch *filter.FilterType {
			case "set":
				if len(*filter.Values) > 0 {
					// проверка на наличие в set null значения
					var nullExistInSet = false
					for _, val := range *filter.Values {
						if val == "" {
							nullExistInSet = true
							break
						}
					}
					if nullExistInSet { // если null есть - то добавляем условие OR _ IS NULL
						tbl = tbl.Where(fmt.Sprintf("%s in (?) OR %s IS NULL", field, field), *filter.Values)
					} else {
						tbl = tbl.Where(fmt.Sprintf("%s in (?)", field), *filter.Values)
					}
				}
			case "date":
				if filterOperator == "" {
					tbl = constructDateWhere(tbl, field, filterOperator, filter)
				} else {
					tbl = constructDateWhere(tbl, field, filterOperator, filter.Condition1.Filter, filter.Condition2.Filter)
				}
			case "number":
				tbl = constructNumberWhere(tbl, field, filter)
			case "text":
				if filterOperator == "" {
					tbl = constructTextWhere(tbl, field, filterOperator, filter)
				} else {
					tbl = constructTextWhere(tbl, field, filterOperator, filter.Condition1.Filter, filter.Condition2.Filter)
				}
			default:
				log.Println("unknown number filterType: " + *filter.FilterType)
				return tbl
			}
		}
	}
	return tbl
}

func constructNumberWhere(tbl *bun.SelectQuery, field string, filter models.Filter) *bun.SelectQuery {
	operators := map[string]string{
		"equals":             "%s = ?",
		"notEqual":           "%s <> ?",
		"greaterThan":        "%s > ?",
		"greaterThanOrEqual": "%s >= ?",
		"lessThan":           "%s < ?",
		"lessThanOrEqual":    "%s <= ?",
		"inRange":            "%s between ? and ?",
	}
	tbl = constructQuery(tbl, operators[*filter.Type], "", field, "", fmt.Sprintf("%v", *filter.Filter))
	return tbl
}

func constructDateWhere(tbl *bun.SelectQuery, field string, operator string, filters ...models.Filter) *bun.SelectQuery {
	var start1, start2 string
	start1 = getTimeString(*filters[0].DateFrom)
	if operator != "" {
		start2 = getTimeString(*filters[1].DateFrom)
	}
	dateOperators := map[string]string{
		"inRange":     "DATE(%s) between ? and ?",
		"equals":      "DATE(%s) = ?",
		"greaterThan": "DATE(%s) > ?",
		"lessThan":    "DATE(%s) < ?",
		"notEqual":    "DATE(%s) <> ?",
	}
	if *filters[0].Type == "inRange" {
		end1 := getTimeString(*filters[0].DateTo)
		if operator == "" {
			tbl = tbl.Where(fmt.Sprintf("DATE(%s) between ? and ?", field), start1, end1)
		} else {
			end2 := getTimeString(*filters[1].DateTo)
			tbl = tbl.Where(fmt.Sprintf("DATE(%s) between ? and ? %s DATE(%s) between ? and ?", field, operator, field),
				start1, end1, start2, end2)
		}
	} else {
		if operator == "" {
			tbl = constructQuery(tbl, dateOperators[*filters[0].Type], "", field, operator, start1)
		} else {
			tbl = constructQuery(tbl, dateOperators[*filters[0].Type], dateOperators[*filters[1].Type], field, operator, start1, start2)
		}
	}

	return tbl
}

func constructTextWhere(tbl *bun.SelectQuery, field string, operator string, filters ...models.Filter) *bun.SelectQuery {
	operators := map[string]string{
		"equals":      "%s = ?",
		"notEqual":    "%s <> ?",
		"contains":    "%s LIKE ?",
		"notContains": "%s NOT LIKE ?",
		"startsWith":  "%s LIKE ?",
		"endsWith":    "%s LIKE ?",
	}
	if operator == "" {
		tbl = constructQuery(tbl, operators[*filters[0].Type], "", field, operator, likeMix(*filters[0].Type, fmt.Sprintf("%v", *filters[0].Filter)))
	} else {
		tbl = constructQuery(tbl, operators[*filters[0].Type], operators[*filters[1].Type], field, operator, likeMix(*filters[0].Type, (*filters[0].Filter).(string)), likeMix(*filters[1].Type, (*filters[1].Filter).(string)))
	}
	return tbl
}

func likeMix(typeOperator, filter string) (result string) {
	likeOperators := map[string]string{
		"contains":    "%%%s%%",
		"notContains": "%%%s%%",
		"startsWith":  "%s%%",
		"endsWith":    "%%%s",
	}
	likePhrase, ok := likeOperators[typeOperator]
	if ok {
		return fmt.Sprintf(likePhrase, filter)
	}
	return filter
}

func getTimeString(date string) string {
	resultDate := time.Now()
	resultDate, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		resultDate = time.Date(resultDate.Year(), resultDate.Month(), resultDate.Day(), 0, 0, 0, 0, time.UTC)
	}
	location, _ := time.LoadLocation("Europe/Moscow")
	result := resultDate.In(location).Format("2006-01-02")
	return result
}

func constructQuery(tbl *bun.SelectQuery, query0, query1, field, operator string, params ...string) *bun.SelectQuery {
	if operator == "" {
		tbl = tbl.Where(fmt.Sprintf(query0, field), params[0])
	} else {
		query := fmt.Sprintf(query0, field) + " " + operator + " " + fmt.Sprintf(query1, field)
		tbl = tbl.Where(query, params[0], params[1])
	}
	return tbl
}
