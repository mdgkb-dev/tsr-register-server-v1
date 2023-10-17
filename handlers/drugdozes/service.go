package drugdozes

import (
	"context"
	"errors"
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

func (s *Service) CalculateNeeding(c context.Context, opts DrugNeedingOptions) (*models.DrugNeeding, error) {
	drugDoze, err := R.Get(c, opts.DrugDozeID.UUID.String())
	if err != nil {
		return nil, err
	}
	drugRegimen := drugDoze.DrugRegimens.FindDrugRegimen(30, 24)
	if drugRegimen == nil {
		return nil, errors.New("подходящий режим приёма не найдет. проверьте данные")
	}
	variablesMap := map[string]interface{}{string(models.AnthropomethryKeyWeight): opts.Weight, string(models.AnthropomethryKeyHeight): opts.Height}
	quantity := drugRegimen.CalculateNeeding(variablesMap, *opts.Start, *opts.End)

	drugNeeding := models.DrugNeeding{}
	drugNeeding.Init(drugRegimen, uint(quantity), uint(quantity/drugDoze.GetActiveComponentsSum()), "24 гр 12 фл")
	return &drugNeeding, nil
}
