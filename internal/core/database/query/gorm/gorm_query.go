package gorm

import (
	"backend/internal/core/database/query"
	gormutils "backend/internal/core/database/query/gorm/utils"
	"context"
	"strings"

	"gorm.io/gorm"
)

type dbQuery[T any] struct {
	DB *gorm.DB
}

func New[T any](db *query.AppIDB) query.DB[T] {
	return &dbQuery[T]{DB: db.GormDB}
}

func (g *dbQuery[T]) Count(ctx context.Context, whereObj *T) (int64, error) {
	var obj T
	db := WithContext(ctx, g.DB)
	db = db.Model(&obj).Where(whereObj)

	var total int64
	db = db.Count(&total)
	if err := db.Error; err != nil {
		return 0, gormutils.DbError(err)
	}
	return total, nil
}

func (g *dbQuery[T]) List(ctx context.Context, whereObj *T) (*[]T, error) {
	var result []T
	db := WithContext(ctx, g.DB)
	db = db.Model(&result).Where(whereObj)

	db = db.Find(&result)
	if err := db.Error; err != nil {
		return nil, gormutils.DbError(err)
	}
	return &result, nil
}

func (g *dbQuery[T]) CountWhere(ctx context.Context, whereCauses *[]query.WhereCause) (int64, error) {
	var obj T
	db := WithContext(ctx, g.DB)
	db = db.Model(&obj)

	addWhere(db, whereCauses)

	var total int64
	db = db.Count(&total)
	if err := db.Error; err != nil {
		return 0, gormutils.DbError(err)
	}
	return total, nil
}

func (g *dbQuery[T]) ListWhere(ctx context.Context, whereCauses *[]query.WhereCause, limit *query.Limit, sortBy []string) (*[]T, error) {
	var result []T
	db := WithContext(ctx, g.DB)

	addWhere(db, whereCauses)

	if limit != nil {
		db.Limit(int(limit.PageSize)).Offset(int(limit.Offset))
	}
	if sortBy != nil {
		db.Order(strings.Join(sortBy, ","))
	}

	db = db.Find(&result)
	if err := db.Error; err != nil {
		return nil, gormutils.DbError(err)
	}
	return &result, nil
}

func (g *dbQuery[T]) QueryListWhere(ctx context.Context, whereCauses *[]query.WhereCause, limit *query.Limit, sortBy []string) (*[]T, int64, error) {
	var result []T
	db := WithContext(ctx, g.DB)

	addWhere(db, whereCauses)

	var total int64
	db = db.Count(&total)
	if err := db.Error; err != nil {
		return nil, 0, gormutils.DbError(err)
	}

	if limit != nil {
		db.Limit(int(limit.PageSize)).Offset(int(limit.Offset))
	}
	if sortBy != nil {
		db.Order(strings.Join(sortBy, ","))
	}

	db = db.Find(&result)
	if err := db.Error; err != nil {
		return nil, 0, gormutils.DbError(err)
	}
	return &result, total, nil
}

func (g *dbQuery[T]) RawQuery(ctx context.Context, sqlQuery string, params []interface{}, result interface{}) error {
	db := WithContext(ctx, g.DB)
	db.Raw(sqlQuery, params...).Scan(result)
	if err := db.Error; err != nil {
		return gormutils.DbError(err)
	}
	return nil
}

func (g *dbQuery[T]) RawExec(ctx context.Context, sqlQuery string, params []interface{}) (int64, error) {
	db := WithContext(ctx, g.DB)
	dbe := db.Exec(sqlQuery, params...)
	if err := dbe.Error; err != nil {
		return 0, gormutils.DbError(err)
	}
	return dbe.RowsAffected, nil
}

func (g *dbQuery[T]) Create(ctx context.Context, obj *T) (*T, error) {
	db := WithContext(ctx, g.DB)
	db = db.Create(&obj)
	if err := db.Error; err != nil {
		return nil, gormutils.DbError(err)
	}
	return obj, nil
}

func (g *dbQuery[T]) CreateList(ctx context.Context, obj *[]T) (*[]T, error) {
	db := WithContext(ctx, g.DB)
	db = db.Create(&obj)
	if err := db.Error; err != nil {
		return nil, gormutils.DbError(err)
	}
	return obj, nil
}

func (g *dbQuery[T]) Update(ctx context.Context, obj *T) (*T, error) {
	db := WithContext(ctx, g.DB)
	db = db.Updates(obj)
	if err := db.Error; err != nil {
		return nil, gormutils.DbError(err)
	}
	return obj, nil
}

func (g *dbQuery[T]) UpdateList(ctx context.Context, obj *[]T) (*[]T, error) {
	for _, row := range *obj {
		o := row
		_, err := g.Update(ctx, &o)
		if err != nil {
			return nil, gormutils.DbError(err)
		}
	}
	return obj, nil
}

func (g *dbQuery[T]) UpdateForce(ctx context.Context, obj *T) (*T, error) {
	db := WithContext(ctx, g.DB)
	db = db.Select("*").Updates(obj)
	if err := db.Error; err != nil {
		return nil, gormutils.DbError(err)
	}
	return obj, nil
}

func (g *dbQuery[T]) FindByPK(ctx context.Context, obj *T) (*T, error) {
	var result T
	db := WithContext(ctx, g.DB)
	if err := db.Where(obj).First(&result).Error; err != nil {
		return nil, gormutils.DbError(err)
	}
	return &result, nil
}

func (g *dbQuery[T]) FirstWhere(ctx context.Context, whereCauses *[]query.WhereCause, sortBy []string) (*T, error) {
	var result T
	db := WithContext(ctx, g.DB)

	addWhere(db, whereCauses)

	if sortBy != nil {
		db.Order(strings.Join(sortBy, ","))
	}

	db = db.First(&result)
	if err := db.Error; err != nil {
		return nil, gormutils.DbError(err)
	}
	return &result, nil
}

func (g *dbQuery[T]) FindOneWhere(ctx context.Context, whereCauses *[]query.WhereCause) (*T, error) {
	var result T
	db := WithContext(ctx, g.DB)

	addWhere(db, whereCauses)

	db = db.Find(&result)
	if err := db.Error; err != nil {
		return nil, gormutils.DbError(err)
	}
	return &result, nil
}

func (g *dbQuery[T]) FindOne(ctx context.Context, obj *T) (*T, error) {
	var result T
	db := WithContext(ctx, g.DB)
	if err := db.Where(obj).First(&result).Error; err != nil {
		return nil, gormutils.DbError(err)
	}
	return &result, nil
}

func (g *dbQuery[T]) Delete(ctx context.Context, obj *T) error {
	db := WithContext(ctx, g.DB)
	db = db.Delete(obj)
	if err := db.Error; err != nil {
		return gormutils.DbError(err)
	}
	return nil
}

func (g *dbQuery[T]) DeleteList(ctx context.Context, obj *[]T) error {
	db := WithContext(ctx, g.DB)
	for _, o := range *obj {
		db = db.Delete(o)
		if err := db.Error; err != nil {
			return gormutils.DbError(err)
		}
	}
	return nil
}

func (g *dbQuery[T]) DeleteWhere(ctx context.Context, whereCauses *[]query.WhereCause) error {
	var obj T
	db := WithContext(ctx, g.DB)

	addWhere(db, whereCauses)

	db = db.Delete(&obj)
	if err := db.Error; err != nil {
		return gormutils.DbError(err)
	}
	return nil
}

func WithContext(ctx context.Context, db *gorm.DB) *gorm.DB {
	return db.WithContext(ctx)
}

func addWhere(db *gorm.DB, whereCauses *[]query.WhereCause) {
	if whereCauses != nil {
		for _, w := range *whereCauses {
			if w.Type == query.And {
				db.Where(w.Query, w.Args...)
			} else {
				db.Or(w.Query, w.Args...)
			}
		}
	}
}
