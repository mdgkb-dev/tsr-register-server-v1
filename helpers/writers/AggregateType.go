package writers

type AggregateType string

const (
	AggregateNone     AggregateType = "none"
	AggregateSum      AggregateType = "sum"
	AggregateExisting AggregateType = "existing"
	AggregateAverage  AggregateType = "average"
)

const (
	Yes    string = "Да"
	No     string = "Нет"
	NoData string = "Нет данных"
)

// func (item AggregateType) GetAggregatedPercentage(aggregatedValues map[string]float64) ResearchQueryPercentages {
// 	r := make(ResearchQueryPercentages, 0)
// 	sum := float64(0)
// 	for _, v := range aggregatedValues {
// 		sum += v
// 	}
// 	if item == AggregateNone || item == "" || item == AggregateExisting {
// 		for k, v := range aggregatedValues {
// 			perc := v * 100 / sum
// 			r = append(r, &ResearchQueryPercentage{k, perc})
// 		}
// 		sort.Slice(r, func(i, j int) bool {
// 			return r[i].Value > r[j].Value
// 		})
// 	}
// 	if item == AggregateAverage {
// 		r = append(r, &ResearchQueryPercentage{Key: "Среднее значение", Value: sum / float64(len(aggregatedValues))})
// 	}
// 	return
// }

// func (item AggregateType) WriteAggregatedValues(xl *writers.XlsxHelper, aggregatedValues map[string]float64) {
// 	for i, data := range item.GetAggregatedPercentage(aggregatedValues) {
// 		d := fmt.Sprintf("%s : %.2f%%", data.Key, data.Value)
// 		xl.WriteString(xl.StrCursor+i, xl.Cursor, &[]string{d})
// 	}
// 	xl.Cursor++
// }
