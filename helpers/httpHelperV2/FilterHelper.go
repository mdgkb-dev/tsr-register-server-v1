package httpHelper

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryFilter struct {
	ID           *string
	UserID uuid.UUID
	FilterModels FilterModels
	SortModels   SortModels
	Pagination   *Pagination
	WithDeleted  bool
}

type Pagination struct {
	Offset *int
	Limit  *int
}

func (i *HTTPHelper) CreateQueryFilter(c *gin.Context) (*QueryFilter, error) {
	filterModels, err := createFilterModels(c)
	if err != nil {
		return nil, err
	}
	sortModels, err := createSortModels(c)
	if err != nil {
		return nil, err
	}
	pagination, err := createPagination(c)
	if err != nil {
		return nil, err
	}
	id := c.Param("id")
	userID, err := models.GetUserID(c)
	if err != nil {
		return nil, err
	}
	return &QueryFilter{ID: &id, UserID: *userID, FilterModels: filterModels, SortModels: sortModels, Pagination: pagination}, nil
}

func createSortModels(c *gin.Context) (SortModels, error) {
	sortModels := make(SortModels, 0)
	if c.Query("sortModel") == "" {
		return nil, nil
	}
	for _, arg := range c.QueryArray("sortModel") {
		sortModel, err := ParseJSONToSortModel(arg)
		if err != nil {
			return nil, err
		}
		sortModels = append(sortModels, &sortModel)
	}

	return sortModels, nil
}

func createFilterModels(c *gin.Context) (FilterModels, error) {
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

func createPagination(c *gin.Context) (*Pagination, error) {

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
	fmt.Println(offsetNumber, limit)
	fmt.Println(offsetNumber, limit)
	fmt.Println(offsetNumber, limit)
	fmt.Println(offsetNumber, limit)
	return &Pagination{Offset: &offsetNumber, Limit: &limit}, nil
}

func (i *HTTPHelper) CreateWithDeleted(c *gin.Context) (bool, error) {
	withDeleted := c.Query("withDeleted")
	if withDeleted == "true" {
		return true, nil
	}
	return false, nil
}

func (i *HTTPHelper) CreateWithDeletedQuery(query *bun.SelectQuery, withDeleted bool) {
	if withDeleted {
		query.WhereAllWithDeleted()
	}
}
