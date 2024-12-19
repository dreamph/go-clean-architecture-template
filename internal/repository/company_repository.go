package repository

import (
	"backend/internal/core/models"
	"backend/internal/core/utils"
	"backend/internal/domain"
	"backend/internal/domain/repomodels"
	"context"
	"database/sql"
	"strings"

	"github.com/dreamph/dbre"
	"github.com/dreamph/dbre/query"
	"github.com/dreamph/dbre/query/bun"
	"github.com/guregu/null"
)

type CompanyRepository interface {
	WithTx(db *query.AppIDB) CompanyRepository
	Create(ctx context.Context, obj *domain.Company) (*domain.Company, error)
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*domain.Company, error)
	FindByCode(ctx context.Context, code string) (*domain.Company, error)
	Update(ctx context.Context, obj *domain.Company) (*domain.Company, error)
	List(ctx context.Context, obj *repomodels.CompanyListCriteria) (*[]domain.Company, int64, error)
	ListData(ctx context.Context, obj *repomodels.CompanyListCriteria) (*[]repomodels.CompanyInfoData, int64, error)
	TestList(ctx context.Context) error
}

type companyRepository struct {
	query query.DB[domain.Company]
}

func NewCompanyRepository(db *query.AppIDB) CompanyRepository {
	return &companyRepository{
		query: bun.New[domain.Company](db),
	}
}

func (r *companyRepository) WithTx(tx *query.AppIDB) CompanyRepository {
	return NewCompanyRepository(tx)
}

func (r *companyRepository) Create(ctx context.Context, obj *domain.Company) (*domain.Company, error) {
	return r.query.Create(ctx, obj)
}

func (r *companyRepository) Delete(ctx context.Context, id string) error {
	return r.query.Delete(ctx, &domain.Company{ID: id})
}

func (r *companyRepository) FindByID(ctx context.Context, id string) (*domain.Company, error) {
	return r.query.FindByPK(ctx, &domain.Company{ID: id})
}

func (r *companyRepository) FindByCode(ctx context.Context, code string) (*domain.Company, error) {
	return r.query.FindOne(ctx, &domain.Company{Code: null.StringFrom(code)})
}

func (r *companyRepository) List(ctx context.Context, obj *repomodels.CompanyListCriteria) (*[]domain.Company, int64, error) {
	result := &[]domain.Company{}
	whereBuilder := query.NewWhereBuilder()

	if obj.Status != 0 {
		whereBuilder.Where("status = ?", obj.Status)
	}

	whereCauses := whereBuilder.WhereCauses()
	total, err := r.query.CountWhere(ctx, whereCauses)
	if err != nil {
		return nil, 0, err
	}
	if total > 0 {
		result, err = r.query.ListWhere(ctx, whereCauses, utils.ToQueryLimit(obj.Limit), []string{})
		if err != nil {
			return nil, 0, err
		}
	}

	return result, total, nil
}

func (r *companyRepository) ListData(ctx context.Context, obj *repomodels.CompanyListCriteria) (*[]repomodels.CompanyInfoData, int64, error) {
	result := &[]repomodels.CompanyInfoData{}
	var total int64

	queryBuilder := dbre.NewSQLQueryBuilder()
	queryBuilder.AddQuery("SELECT c.*, 'TEST' as otherData")
	queryBuilder.AddQuery("FROM company c")
	queryBuilder.AddQuery("WHERE 1 = 1")

	if utils.IsNotEmpty(obj.Code) {
		queryBuilder.AddQueryWithParam("AND c.code LIKE @code",
			sql.Named("code", "%"+obj.Code+"%"),
		)
	}
	if utils.IsNotEmpty(obj.CompanyName) {
		queryBuilder.AddQueryWithParam("AND (LOWER(c.name) LIKE @name OR LOWER(c.name_th) LIKE @nameTh)",
			sql.Named("name", "%"+strings.ToLower(obj.CompanyName)+"%"),
			sql.Named("nameTh", "%"+strings.ToLower(obj.CompanyName)+"%"),
		)
	}
	if obj.Status != 0 {
		queryBuilder.AddQueryWithParam("AND c.status = @status", sql.Named("status", obj.Status))
	}

	mainStatement := queryBuilder.ToSQLQuery()
	countStatement := "select count(1) from (" + mainStatement + ") as t"
	err := r.query.RawQuery(ctx, countStatement, queryBuilder.GetQueryParams(), &total)
	if err != nil {
		return nil, 0, err
	}

	if total > 0 {
		sortSQL, err := dbre.SortSQL(&dbre.SortParam{
			SortFieldMapping: map[string]string{
				"id":           "c.id",
				"code":         "c.code",
				"name":         "c.name",
				"nameTh":       "c.name_th",
				"status":       "c.status",
				"createBy":     "c.create_by",
				"createDate":   "c.create_date",
				"updateBy":     "c.update_by",
				"updateDate":   "c.update_date",
				"createByName": "c.create_by_name",
				"updateByName": "c.update_by_name",
			},
			Sort: obj.Sort,
			DefaultSort: &dbre.Sort{
				SortBy:        "updateDate",
				SortDirection: dbre.DESC,
			},
		})
		if err != nil {
			return nil, 0, err
		}

		if utils.IsNotEmpty(sortSQL) {
			queryBuilder.AddQuery("ORDER BY " + sortSQL)
		}

		pageQuery := utils.GetPageQuery(obj.Limit)
		queryBuilder.AddQueryWithParam("LIMIT @pageSize OFFSET @offset",
			sql.Named("pageSize", pageQuery.PageSize),
			sql.Named("offset", pageQuery.Offset),
		)

		err = r.query.RawQuery(ctx, queryBuilder.ToSQLQuery(), queryBuilder.GetQueryParams(), result)
		if err != nil {
			return nil, 0, err
		}
	}

	return result, total, nil
}

func (r *companyRepository) Update(ctx context.Context, obj *domain.Company) (*domain.Company, error) {
	return r.query.Update(ctx, obj)
}

func (r *companyRepository) TestList(ctx context.Context) error {
	_, err := r.query.ListWhere(ctx, nil, utils.ToQueryLimit(models.LimitForCount), nil)
	if err != nil {
		return err
	}
	return nil
}
