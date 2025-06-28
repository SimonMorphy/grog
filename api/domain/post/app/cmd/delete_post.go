package cmd

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/repo"
	"github.com/SimonMorphy/grog/api/infra/pkg/decorator"
)

type DeletePost struct {
	ID uint
}
type DeletePostResult struct {
}
type DeletePostHandler decorator.Handler[DeletePost, *DeletePostResult]

type deletePostHandler struct {
	Repo repo.PostRepository
}

func (d deletePostHandler) Handle(ctx context.Context, query DeletePost) (*DeletePostResult, error) {
	return nil, d.Repo.Delete(ctx, query.ID)
}
