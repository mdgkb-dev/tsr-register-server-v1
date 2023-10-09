package drugdozes

import (
	"context"
	"errors"
	"fmt"
	"math"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(c context.Context, item *models.DrugDoze) error {
	err := R.Create(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(c context.Context, item *models.DrugDoze) error {
	err := R.Update(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(c context.Context) (models.DrugDozesWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, slug string) (*models.DrugDoze, error) {
	item, err := R.Get(c, slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}

func (s *Service) CalculateNeeding(c context.Context, opts DrugNeedingOptions) (float64, error) {
	drugDoze, err := R.Get(c, opts.DrugDozeID.UUID.String())
	fmt.Println("dd1", drugDoze, err)
	if err != nil {
		return math.NaN(), err
	}
	fmt.Println("dd", drugDoze)
	drugRegimen := drugDoze.DrugRegimens.FindDrugRegimen(30, 24)
	fmt.Println("dr", drugRegimen)
	if drugRegimen == nil {
		return math.NaN(), errors.New("подходящий режим приёма не найдет. проверьте данные")
	}
	variablesMap := map[string]interface{}{string(models.AnthropomethryKeyWeight): opts.Weight, string(models.AnthropomethryKeyHeight): opts.Height}
	fmt.Println("opts", opts)
	quantity := drugRegimen.CalculateNeeding(variablesMap, *opts.Start, *opts.End)

	return quantity / drugDoze.GetActiveComponentsSum(), nil
}
