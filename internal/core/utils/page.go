package utils

import (
	"backend/internal/core/models"

	"github.com/dreamph/dbre"
)

// GetPageResult ...
func GetPageResult(pageLimit *models.PageLimit, total int64) *models.PageResult {
	pageResult := &models.PageResult{}
	pageResult.PageNumber = pageLimit.PageNumber
	pageResult.PageSize = pageLimit.PageSize
	pageResult.Total = total
	totalPages := total / pageLimit.PageSize
	if total%pageLimit.PageSize != 0 {
		totalPages++
	}
	pageResult.TotalPages = totalPages
	return pageResult
}

// GetPageQuery ...
func GetPageQuery(pageLimit *models.PageLimit) *models.PageQuery {
	pageResult := &models.PageQuery{}
	pageResult.PageSize = pageLimit.PageSize
	pageResult.Offset = (pageLimit.PageNumber - 1) * pageLimit.PageSize
	return pageResult
}

// ToQueryLimit ...
func ToQueryLimit(pageLimit *models.PageLimit) *dbre.Limit {
	if pageLimit == nil {
		return nil
	}
	limit := &dbre.Limit{}
	limit.PageSize = pageLimit.PageSize
	limit.Offset = (pageLimit.PageNumber - 1) * pageLimit.PageSize
	return limit
}

// GetPageQueryByPageNumber ...
func GetPageQueryByPageNumber(pageNumber int64, pageSize int64) *models.PageQuery {
	pageResult := &models.PageQuery{}
	pageResult.PageSize = pageSize
	pageResult.Offset = (pageNumber - 1) * pageSize
	return pageResult
}
