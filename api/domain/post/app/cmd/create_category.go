package cmd

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/entity"
	"github.com/SimonMorphy/grog/api/domain/post/repo"
	"github.com/SimonMorphy/grog/api/infra/pkg/decorator"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CreateCategory struct {
	Name string
}

type CreateCategoryResult struct {
	Category *entity.Category
}

type CreateCategoryHandler decorator.Handler[CreateCategory, *CreateCategoryResult]

type createCategoryHandler struct {
	Repo repo.CategoryRepository
}

func (c createCategoryHandler) Handle(ctx context.Context, query CreateCategory) (*CreateCategoryResult, error) {
	res, err := c.Repo.Create(ctx, &entity.Category{Name: query.Name})
	if err != nil {
		logrus.Errorf(
			"failed to create category: %v, error: %v", query, err)
		return nil, err
	}
	return &CreateCategoryResult{
		Category: res,
	}, nil
}

func NewCreateCategoryHandler(
	repo repo.CategoryRepository,
	entry *logrus.Entry,
	record decorator.MetricsRecord,
) CreateCategoryHandler {
	if repo == nil {
		logrus.Panic(gorm.ErrInvalidDB)
	}
	return decorator.ApplyHandlerDecorators[CreateCategory, *CreateCategoryResult](
		&createCategoryHandler{Repo: repo},
		entry,
		record,
	)
}
