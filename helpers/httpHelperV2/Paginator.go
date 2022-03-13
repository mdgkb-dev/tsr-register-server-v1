package httpHelper

import (
	"github.com/uptrace/bun"
)

func (i *HTTPHelper) CreatePaginationQuery(query *bun.SelectQuery, pagination *Pagination) {
	if pagination == nil {
		return
	}
	query = query.Offset(*pagination.Offset)
	query = query.Limit(*pagination.Limit)
}
