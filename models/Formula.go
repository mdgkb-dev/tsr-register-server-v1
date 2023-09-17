package models

import (
	"math"

	"github.com/Pramod-Devireddy/go-exprtk"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Formula struct {
	bun.BaseModel `bun:"formulas,alias:formulas"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string        `json:"name"`
	Formula       string        `json:"formula"`
	Color         string        `json:"color"`
	Xlsx          bool          `json:"xlsx"`
	Research      *Research     `bun:"rel:belongs-to" json:"research"`
	ResearchID    uuid.NullUUID `bun:"type:uuid" json:"researchId"`

	FormulaResults          FormulaResults `bun:"rel:has-many" json:"formulaResults"`
	FormulaResultsForDelete []uuid.UUID    `bun:"-" json:"formulaResultsForDelete"`
}

type Formulas []*Formula

func (items Formulas) SetXlsxData(variables map[string]interface{}) ([]interface{}, error) {
	results := make([]interface{}, 0)
	m := exprtk.NewExprtk()
	defer m.Delete()
	for i := range items {
		formulaResult, err := items[i].SetXlsxData(variables, m)
		if err != nil {
			break
		}
		results = append(results, formulaResult...)
	}
	return results, nil
}

func (item *Formula) SetXlsxData(variables map[string]interface{}, m exprtk.GoExprtk) ([]interface{}, error) {
	results := make([]interface{}, 0)
	if !item.Xlsx {
		return results, nil
	}
	m.SetExpression(item.Formula)
	for k := range variables {
		m.AddDoubleVariable(k)
	}
	err := m.CompileExpression()
	if err != nil {
		return results, err
	}
	for k, v := range variables {
		switch v := v.(type) {
		case float32:
			m.SetDoubleVariableValue(k, float64(v))
		case float64:
			m.SetDoubleVariableValue(k, v)
		case int:
			m.SetDoubleVariableValue(k, float64(int64(v)))
		}
	}
	value := m.GetEvaluatedValue()
	if math.IsNaN(value) {
		results = append(results, "Ошибка в рассчёте")
	} else {
		results = append(results, value)
	}
		

	if len(item.FormulaResults) > 0 {
		result := item.GetResult(value)
		if result != nil {
			results = append(results, result.Name)
		} else {
			results = append(results, "")
		}
	}

	return results, nil
}

func (item *Formula) GetResult(value float64) (result *FormulaResult) {
	for _, formulaResult := range item.FormulaResults {
		if value > formulaResult.LowRange && value < formulaResult.HighRange {
			result = formulaResult
			break
		}
	}
	return result
}
