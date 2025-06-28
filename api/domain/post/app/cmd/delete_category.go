package cmd

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/repo"
	"github.com/SimonMorphy/grog/api/infra/pkg/decorator"
)

type DeleteCategory struct {
	ID uint
}

type DeleteCategoryResult struct {
}

type DeleteCategoryHandler decorator.CommandHandler[DeleteCategory, *DeleteCategoryResult]

type deleteCategoryHandler struct {
	Repo repo.CategoryRepository
}

func (c deleteCategoryHandler) Handle(ctx context.Context, cmd DeleteCategory) (*DeleteCategoryResult, error) {
	if err := c.Repo.Delete(ctx, cmd.ID); err != nil {
		return nil, err
	}
	return &DeleteCategoryResult{}, nil
}
