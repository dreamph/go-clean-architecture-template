package utils

import (
	"backend/internal/core/models"
)

/*
QueryAll

	//ex1
	limit := &models.PageLimit{PageNumber: 1, PageSize: 10}
	criteria := &repomodels.DigitalCertificateListCriteria{
		Limit: limit,
	}
	_, total, err := utils.QueryAll(limit, func(currentLimit *models.PageLimit) (*[]repomodels.DigitalCertificateData, int64, error) {
		criteria.Limit = currentLimit
		return repositoryRegistry.DigitalCertificateRepository.List(ctx, criteria)
	})

	//ex2
	limit := &models.PageLimit{PageNumber: 1, PageSize: 3}
	criteria := &repomodels.UserListCriteria{
		Limit: limit,
	}
	listUser, total, err := utils.QueryAll(limit, func(currentLimit *models.PageLimit) (*[]domain.User, int64, error) {
		return repositoryRegistry.UserRepository.List(ctx, criteria)
	})

	//ex3
	criteria := &repomodels.UserListCriteria{
		Limit: &models.PageLimit{PageNumber: 1, PageSize: 3},
	}
	listUser, total, err := utils.QueryAll(criteria.Limit, func(currentLimit *models.PageLimit) (*[]domain.User, int64, error) {
		return repositoryRegistry.UserRepository.List(ctx, criteria)
	})
*/
func QueryAll[T any](limit *models.PageLimit, repoQueryFunction func(limit *models.PageLimit) (*[]T, int64, error)) (*[]T, int64, error) {
	var list []T
	data, total, err := repoQueryFunction(limit)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, nil
	}
	list = append(list, *data...)
	limitResult := GetPageResult(limit, total)
	if limitResult.TotalPages == 1 {
		return &list, total, nil
	}

	totalPages := limitResult.TotalPages
	for i := int64(2); i <= totalPages; i++ {
		limit.PageNumber = i

		data, total, err := repoQueryFunction(limit)
		if err != nil {
			return nil, 0, err
		}
		if total == 0 {
			break
		}
		if data == nil {
			break
		}
		list = append(list, *data...)
	}
	return &list, int64(len(list)), nil
}

/*
	criteria := &repomodels.UserListCriteria{
		Limit: &models.PageLimit{PageNumber: 1, PageSize: 3},
	}
	listUser, total, err := utils.QueryAll(criteria.Limit, func(currentLimit *models.PageLimit) (*[]domain.User, int64, error) {
		return repositoryRegistry.UserRepository.List(ctx, criteria)
	})

	fmt.Println("======")
	fmt.Println(total)
	for i, u := range *listUser {
		fmt.Println(i)
		fmt.Println(u.Email)
	}
	fmt.Println("======")
*/

func QueryIn[T any, R any](data []T, size int, f func(values []T) ([]R, error)) (*[]R, error) {
	chunkedList := Chunk(data, size)
	var result []R
	for i := range chunkedList {
		list, err := f(chunkedList[i])
		if err != nil {
			return nil, err
		}
		result = append(result, list...)
	}
	return &result, nil
}
