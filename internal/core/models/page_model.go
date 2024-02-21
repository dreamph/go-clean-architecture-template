package models

// PageLimit ...
type PageLimit struct {
	PageNumber int64 `json:"pageNumber" example:"1"`
	PageSize   int64 `json:"pageSize" example:"10"`
}

type Sort struct {
	SortBy        string `json:"sortBy"`
	SortDirection string `json:"sortDirection" example:"DESC" enums:"ASC,DESC"`
}

// PageResult ...
type PageResult struct {
	PageSize   int64 `json:"pageSize"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"totalPages"`
	PageNumber int64 `json:"pageNumber"`
}

type PageQuery struct {
	Offset   int64 `json:"offset"`
	PageSize int64 `json:"pageSize"`
}

var (
	MaxLimitForQuery = &PageLimit{
		PageNumber: 1,
		PageSize:   1000,
	}

	MaxLimitForLargeQuery = &PageLimit{
		PageNumber: 1,
		PageSize:   5000,
	}

	LimitForCount = &PageLimit{
		PageNumber: 1,
		PageSize:   1,
	}
)
