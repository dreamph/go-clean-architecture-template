package query

import (
	"context"

	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

type DB[T any] interface {
	RawExec(ctx context.Context, sqlQuery string, params []interface{}) (int64, error)
	RawQuery(ctx context.Context, sql string, params []interface{}, resultPtr interface{}) error
	Create(ctx context.Context, obj *T) (*T, error)
	CreateList(ctx context.Context, obj *[]T) (*[]T, error)
	Update(ctx context.Context, obj *T) (*T, error)
	UpdateList(ctx context.Context, obj *[]T) (*[]T, error)
	UpdateForce(ctx context.Context, obj *T) (*T, error)
	FindByPK(ctx context.Context, obj *T) (*T, error)
	Delete(ctx context.Context, obj *T) error
	DeleteList(ctx context.Context, obj *[]T) error
	DeleteWhere(ctx context.Context, whereCauses *[]WhereCause) error
	Count(ctx context.Context, whereObj *T) (int64, error)
	List(ctx context.Context, whereObj *T) (*[]T, error)
	ListWhere(ctx context.Context, whereCauses *[]WhereCause, limit *Limit, sortBy []string) (*[]T, error)
	QueryListWhere(ctx context.Context, whereCauses *[]WhereCause, limit *Limit, sortBy []string) (*[]T, int64, error)
	FirstWhere(ctx context.Context, whereCauses *[]WhereCause, sortBy []string) (*T, error)
	FindOneWhere(ctx context.Context, whereCauses *[]WhereCause) (*T, error)
	FindOne(ctx context.Context, whereObj *T) (*T, error)
	CountWhere(ctx context.Context, whereCauses *[]WhereCause) (int64, error)
}

type TxFn func(ctx context.Context, appDB *AppIDB) error

type AppIDB struct {
	BunDB  bun.IDB
	GormDB *gorm.DB
}

type DBTx interface {
	WithTx(ctx context.Context, fn TxFn) error
}
