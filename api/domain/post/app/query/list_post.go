package query

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/entity"
	"github.com/SimonMorphy/grog/api/domain/post/repo"
	"github.com/SimonMorphy/grog/api/infra/pkg/decorator"
	"github.com/sirupsen/logrus"
)

type ListPost struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type ListPostResult struct {
	Posts []*entity.Post
}

type ListPostHandler decorator.QueryHandler[ListPost, *ListPostResult]

type listPostHandler struct {
	Repo repo.PostRepository
}

func (l listPostHandler) Handle(ctx context.Context, lp ListPost) (*ListPostResult, error) {
	posts, err := l.Repo.List(ctx, lp.Page, lp.Size)
	if err != nil {
		return nil, err
	}
	return &ListPostResult{posts}, nil
}

func NewListPostHandler(
	repo repo.PostRepository,
	entry *logrus.Entry,
	recoder decorator.MetricsRecord,
) ListPostHandler {
	if repo == nil {
		logrus.Panicf("repo cannot be nil,err -> %v", repo)
	}
	return decorator.ApplyHandlerDecorators[ListPost, *ListPostResult](
		&listPostHandler{Repo: repo},
		entry,
		recoder,
	)
}
