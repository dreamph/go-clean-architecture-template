package database

import (
	cerrors "backend/internal/core/errors"
	"backend/internal/core/models"
	"backend/internal/core/utils"

	errs "github.com/pkg/errors"

	"fmt"
)

const (
	ASC  = "ASC"
	DESC = "DESC"
)

type SortParam struct {
	SortFieldMapping map[string]string
	Sort             *models.Sort
	DefaultSort      *models.Sort
}

func SortSQL(param *SortParam) (string, error) {
	if param == nil {
		return "", errs.New("required sortParam")
	}
	if param.SortFieldMapping == nil {
		return "", errs.Wrap(cerrors.ErrValidationFailed, "required SortFieldMapping")
	}

	if param.Sort == nil && param.DefaultSort == nil {
		return "", nil
	}

	if param.Sort == nil {
		return fmt.Sprintf("%s %s", param.SortFieldMapping[param.DefaultSort.SortBy], param.DefaultSort.SortDirection), nil
	}

	if utils.IsEmpty(param.Sort.SortBy) {
		return fmt.Sprintf("%s %s", param.SortFieldMapping[param.DefaultSort.SortBy], param.DefaultSort.SortDirection), nil
	}

	dbField, ok := param.SortFieldMapping[param.Sort.SortBy]
	if !ok {
		return "", errs.Wrap(cerrors.ErrValidationFailed, "sortBy not support :"+param.Sort.SortBy)
	}

	return fmt.Sprintf("%s %s", dbField, param.Sort.SortDirection), nil
}
