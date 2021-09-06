package httpHelper

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryFilter struct {
	FilterModels FilterModels
	Pagination   *Pagination
}

type Pagination struct {
	Offset *int
	Limit  *int
}

func CreateQueryFilter(c *gin.Context) (*QueryFilter, error) {
	filterModels, err := CreateFilterModels(c)
	fmt.Println(filterModels)
	if err != nil {
		return nil, err
	}
	pagination, err := CreatePagination(c)
	if err != nil {
		return nil, err
	}
	return &QueryFilter{FilterModels: filterModels, Pagination: pagination}, nil
}

func CreateFilterModels(c *gin.Context) (FilterModels, error) {
	filterModels := make(FilterModels, 0)
	if c.Query("filterModel") == "" {
		return nil, nil
	}
	for _, arg := range c.QueryArray("filterModel") {
		filterModel, err := ParseJSONToFilterModel(arg)
		if err != nil {
			return nil, err
		}
		filterModels = append(filterModels, &filterModel)
	}

	return filterModels, nil
}

func CreatePagination(c *gin.Context) (*Pagination, error) {
	offset := c.Query("offset")
	if offset == "" {
		return nil, nil
	}
	offsetNumber, err := strconv.Atoi(offset)
	if err != nil {
		return nil, err
	}
	offsetNumber = offsetNumber * 25
	limit := 25
	return &Pagination{Offset: &offsetNumber, Limit: &limit}, nil
}
