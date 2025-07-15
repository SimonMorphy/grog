package cmd

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/repo"
	"github.com/SimonMorphy/grog/api/infra/pkg/decorator"
	"github.com/sirupsen/logrus"
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

func NewDeletePostHandler(
	repo repo.PostRepository,
	entry *logrus.Entry,
	record decorator.MetricsRecord,
) DeletePostHandler {
	if record == nil {
		logrus.Panic("category record is nil")
	}
	return decorator.ApplyHandlerDecorators[DeletePost, *DeletePostResult](
		&deletePostHandler{Repo: repo},
		entry,
		record,
	)
}
