package query

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/entity"
	"github.com/SimonMorphy/grog/api/domain/post/repo"
	"github.com/SimonMorphy/grog/api/infra/pkg/decorator"
)

type GetCategory struct {
	ID uint
}

type GetCategoryResult struct {
	Category *entity.Category
}

type GetCategoryHandler decorator.QueryHandler[GetCategory, *GetCategoryResult]

type getCategoryHandler struct {
	Repo repo.CategoryRepository
}

func (g getCategoryHandler) Handle(ctx context.Context, query GetCategory) (*GetCategoryResult, error) {
	res, err := g.Repo.Get(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	return &GetCategoryResult{Category: res}, nil
}
