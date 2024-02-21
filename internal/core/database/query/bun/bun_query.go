package bun

import (
	"backend/internal/core/database/query"
	bunutils "backend/internal/core/database/query/bun/utils"
	"backend/internal/core/utils"
	"context"

	"github.com/uptrace/bun"
)

type dbQuery[T any] struct {
	DB bun.IDB
}

func New[T any](db *query.AppIDB) query.DB[T] {
	return &dbQuery[T]{DB: db.BunDB}
}

func (q *dbQuery[T]) Count(ctx context.Context, whereObj *T) (int64, error) {
	db := q.DB.NewSelect().Model((*T)(nil))

	q.addWhere(db, bunutils.BuildWhereCause(whereObj))

	total, err := db.Count(ctx)
	if err != nil {
		return 0, bunutils.DbError(err)
	}
	return int64(total), nil
}

func (q *dbQuery[T]) List(ctx context.Context, whereObj *T) (*[]T, error) {
	var result []T
	db := q.DB.NewSelect().Model(&result)

	q.addWhere(db, bunutils.BuildWhereCause(whereObj))

	err := db.Scan(ctx)
	if err != nil {
		return nil, bunutils.DbError(err)
	}
	return &result, nil
}

func (q *dbQuery[T]) CountWhere(ctx context.Context, whereCauses *[]query.WhereCause) (int64, error) {
	db := q.DB.NewSelect().Model((*T)(nil))

	q.addWhere(db, whereCauses)

	total, err := db.Count(ctx)
	if err != nil {
		return 0, bunutils.DbError(err)
	}
	return int64(total), nil
}

func (q *dbQuery[T]) ListWhere(ctx context.Context, whereCauses *[]query.WhereCause, limit *query.Limit, sortBy []string) (*[]T, error) {
	var result []T
	db := q.DB.NewSelect().Model(&result)

	q.addWhere(db, whereCauses)

	if limit != nil {
		db.Limit(int(limit.PageSize)).Offset(int(limit.Offset))
	}
	if sortBy != nil {
		db.Order(sortBy...)
	}

	err := db.Scan(ctx)
	if err != nil {
		return nil, bunutils.DbError(err)
	}
	return &result, nil
}

func (q *dbQuery[T]) QueryListWhere(ctx context.Context, whereCauses *[]query.WhereCause, limit *query.Limit, sortBy []string) (*[]T, int64, error) {
	var result []T
	db := q.DB.NewSelect().Model(&result)

	q.addWhere(db, whereCauses)

	total, err := db.Count(ctx)
	if err != nil {
		return nil, 0, bunutils.DbError(err)
	}

	if limit != nil {
		db.Limit(int(limit.PageSize)).Offset(int(limit.Offset))
	}
	if sortBy != nil {
		db.Order(sortBy...)
	}

	err = db.Scan(ctx)
	if err != nil {
		return nil, 0, bunutils.DbError(err)
	}
	return &result, int64(total), nil
}

func (q *dbQuery[T]) FirstWhere(ctx context.Context, whereCauses *[]query.WhereCause, sortBy []string) (*T, error) {
	var result T
	db := q.DB.NewSelect().Model(&result)

	q.addWhere(db, whereCauses)

	db.Limit(1)

	if sortBy != nil {
		db.Order(sortBy...)
	}

	err := db.Scan(ctx)
	if err != nil {
		return nil, bunutils.DbError(err)
	}
	return &result, nil
}

func (q *dbQuery[T]) FindOneWhere(ctx context.Context, whereCauses *[]query.WhereCause) (*T, error) {
	var result T
	db := q.DB.NewSelect().Model(&result)

	q.addWhere(db, whereCauses)

	err := db.Scan(ctx)
	if err != nil {
		return nil, bunutils.DbError(err)
	}
	return &result, nil
}

func (q *dbQuery[T]) FindOne(ctx context.Context, whereObj *T) (*T, error) {
	var result T
	db := q.DB.NewSelect().Model(&result)

	q.addWhere(db, bunutils.BuildWhereCause(whereObj))

	err := db.Scan(ctx)
	if err != nil {
		return nil, bunutils.DbError(err)
	}
	return &result, nil
}

func (q *dbQuery[T]) RawQuery(ctx context.Context, sqlQuery string, params []interface{}, result interface{}) error {
	namedParameterQuery := bunutils.NewNamedParameterQuery(sqlQuery, params)
	db := q.DB.NewRaw(namedParameterQuery.GetParsedQuery(), namedParameterQuery.GetParsedParameters()...)
	err := db.Scan(ctx, result)
	if err != nil {
		return bunutils.DbError(err)
	}
	return nil
}

func (q *dbQuery[T]) RawExec(ctx context.Context, sqlQuery string, params []interface{}) (int64, error) {
	namedParameterQuery := bunutils.NewNamedParameterQuery(sqlQuery, params)
	db, err := q.DB.ExecContext(ctx, namedParameterQuery.GetParsedQuery(), namedParameterQuery.GetParsedParameters()...)
	if err != nil {
		return 0, bunutils.DbError(err)
	}
	rowsAffected, err := db.RowsAffected()
	if err != nil {
		return 0, bunutils.DbError(err)
	}
	return rowsAffected, nil
}

func (q *dbQuery[T]) Create(ctx context.Context, obj *T) (*T, error) {
	_, err := q.DB.NewInsert().Model(obj).Exec(ctx)
	if err != nil {
		return nil, bunutils.DbError(err)
	}
	return obj, nil
}

func (q *dbQuery[T]) CreateList(ctx context.Context, obj *[]T) (*[]T, error) {
	_, err := q.DB.NewInsert().Model(obj).Exec(ctx)
	if err != nil {
		return nil, bunutils.DbError(err)
	}
	return obj, nil
}

func (q *dbQuery[T]) Update(ctx context.Context, obj *T) (*T, error) {
	_, err := q.DB.NewUpdate().Model(obj).OmitZero().WherePK().Exec(ctx)
	if err != nil {
		return nil, bunutils.DbError(err)
	}
	return obj, nil
}

func (q *dbQuery[T]) UpdateList(ctx context.Context, obj *[]T) (*[]T, error) {
	_, err := q.DB.NewUpdate().Model(obj).OmitZero().Bulk().Exec(ctx)
	if err != nil {
		return nil, bunutils.DbError(err)
	}
	return obj, nil
}

func (q *dbQuery[T]) UpdateForce(ctx context.Context, obj *T) (*T, error) {
	_, err := q.DB.NewUpdate().Model(obj).WherePK().Exec(ctx)
	if err != nil {
		return nil, bunutils.DbError(err)
	}
	return obj, nil
}

func (q *dbQuery[T]) FindByPK(ctx context.Context, obj *T) (*T, error) {
	var result T
	err := q.DB.NewSelect().Model(obj).WherePK().Scan(ctx, &result)
	if err != nil {
		return nil, bunutils.DbError(err)
	}
	return &result, nil
}

func (q *dbQuery[T]) Delete(ctx context.Context, obj *T) error {
	_, err := q.DB.NewDelete().Model(obj).WherePK().Exec(ctx)
	if err != nil {
		return bunutils.DbError(err)
	}
	return nil
}

func (q *dbQuery[T]) DeleteList(ctx context.Context, obj *[]T) error {
	_, err := q.DB.NewDelete().Model(obj).WherePK().Exec(ctx)
	if err != nil {
		return bunutils.DbError(err)
	}
	return nil
}

func (q *dbQuery[T]) DeleteWhere(ctx context.Context, whereCauses *[]query.WhereCause) error {
	db := q.DB.NewDelete().Model((*T)(nil))

	q.addWhereDelete(db, whereCauses)

	_, err := db.Exec(ctx)
	if err != nil {
		return bunutils.DbError(err)
	}
	return nil
}

func (q *dbQuery[T]) addWhere(selectQuery *bun.SelectQuery, whereCauses *[]query.WhereCause) {
	if whereCauses != nil {
		for _, w := range *whereCauses {
			if w.Type == query.And {
				var queryArgs = q.initQueryArgs(w.Args)
				selectQuery.Where(w.Query, queryArgs...)
			} else {
				var queryArgs = q.initQueryArgs(w.Args)
				selectQuery.WhereOr(w.Query, queryArgs...)
			}
		}
	}
}

func (q *dbQuery[T]) addWhereDelete(selectQuery *bun.DeleteQuery, whereCauses *[]query.WhereCause) {
	if whereCauses != nil {
		for _, w := range *whereCauses {
			if w.Type == query.And {
				var queryArgs = q.initQueryArgs(w.Args)
				selectQuery.Where(w.Query, queryArgs...)
			} else {
				var queryArgs = q.initQueryArgs(w.Args)
				selectQuery.WhereOr(w.Query, queryArgs...)
			}
		}
	}
}

func (q *dbQuery[T]) initQueryArgs(args []interface{}) []interface{} {
	var queryArgs []interface{}
	for _, dataArgs := range args {
		if utils.IsSlice(dataArgs) {
			queryArgs = append(queryArgs, bun.In(dataArgs))
		} else {
			queryArgs = append(queryArgs, dataArgs)
		}
	}
	return queryArgs
}
