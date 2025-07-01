package query

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/entity"
	"github.com/SimonMorphy/grog/api/domain/post/repo"
	"github.com/SimonMorphy/grog/api/infra/pkg/decorator"
)

type ListPost struct {
}

type ListPostResult struct {
	Posts []*entity.Post
}

type ListPostHandler decorator.QueryHandler[ListPost, *ListPostResult]

type listPostHandler struct {
	Repo repo.PostRepository
}

func (l listPostHandler) Handle(ctx context.Context, _ ListPost) (*ListPostResult, error) {
	posts, err := l.Repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return &ListPostResult{posts}, nil
}
