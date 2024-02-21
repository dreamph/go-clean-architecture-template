package builder

import (
	"strings"
)

type SQLQueryBuilder struct {
	params  []interface{}
	builder strings.Builder
}

func NewSQLQueryBuilder() SQLQueryBuilder {
	return SQLQueryBuilder{}
}

func (bd *SQLQueryBuilder) AddQuery(sql string) *SQLQueryBuilder {
	bd.builder.WriteString(" " + sql + " ")
	return bd
}

func (bd *SQLQueryBuilder) AddParam(params ...interface{}) *SQLQueryBuilder {
	var length = len(params)
	for i := 0; i < length; i++ {
		bd.params = append(bd.params, params[i])
	}
	return bd
}

func (bd *SQLQueryBuilder) AddQueryWithParam(sql string, params ...interface{}) *SQLQueryBuilder {
	bd.AddQuery(sql)
	bd.AddParam(params...)
	return bd
}

func (bd *SQLQueryBuilder) ToSQLQuery() string {
	return bd.builder.String()
}

func (bd *SQLQueryBuilder) GetQueryParams() []interface{} {
	return bd.params
}
