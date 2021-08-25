package httpHelper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Offset *int
	Limit  *int
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
