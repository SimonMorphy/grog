package cmd

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/entity"
	"github.com/SimonMorphy/grog/api/domain/post/repo"
	"github.com/SimonMorphy/grog/api/infra/pkg/decorator"
)

type BatchSaveCategory struct {
	Categories []*entity.Category
}

type BatchSaveCategoryResult struct {
	Categories []*entity.Category
}

type BatchSaveCategoryHandler decorator.CommandHandler[BatchSaveCategory, *BatchSaveCategoryResult]

type batchSaveCategoryHandler struct {
	Repo repo.CategoryRepository
}

func (b batchSaveCategoryHandler) Handle(ctx context.Context, query BatchSaveCategory) (*BatchSaveCategoryResult, error) {
	err := b.Repo.BatchSave(ctx, query.Categories)
	if err != nil {
		return nil, err
	}
	return &BatchSaveCategoryResult{Categories: query.Categories}, nil
}
