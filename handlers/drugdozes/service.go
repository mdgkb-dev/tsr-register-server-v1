package drugdozes

import (
	"context"
	"fmt"
	"mdgkb/tsr-tegister-server-v1/handlers/drugneedings"
	"mdgkb/tsr-tegister-server-v1/handlers/drugregimens"
	"mdgkb/tsr-tegister-server-v1/handlers/patients"
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
	fmt.Println("drugDOze:", drugDoze)
	patient, err := patients.S.Get(c, opts.PatientID.UUID.String())
	if err != nil {
		return nil, err
	}
	fmt.Println("patient:", patient)
	drugRegimen, err := drugregimens.S.GetByParameters(c, opts.DrugDozeID, patient.Human.GetMonthsFromBirth(), opts.Weight)
	if err != nil {
		return nil, err
	}
	fmt.Println("drugRegimen:", drugRegimen)
	variablesMap := map[string]interface{}{string(models.AnthropomethryKeyWeight): opts.Weight, string(models.AnthropomethryKeyHeight): opts.Height}

	packsNeeding := drugRegimen.CalculateNeeding(variablesMap, uint(opts.End.Sub(*opts.Start).Hours()/24), drugDoze.GetActiveComponentsSum())

	drugNeeding := models.DrugNeeding{}
	// packs := uint( / quantity)
	// fmt.Println("packs", packs)
	drugNeeding.Init(drugRegimen, uint(packsNeeding*float64(drugDoze.Quantity)), uint(packsNeeding))
	drugNeeding.Weight = opts.Weight
	drugNeeding.AgeInMonths = patient.Human.GetMonthsFromBirth()
	err = drugneedings.S.Create(c, &drugNeeding)
	if err != nil {
		return nil, err
	}
	return &drugNeeding, nil
}
