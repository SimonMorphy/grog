package cmd

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/app/dto"
	"github.com/SimonMorphy/grog/api/domain/post/entity"
	"github.com/SimonMorphy/grog/api/domain/post/repo"
	"github.com/SimonMorphy/grog/api/infra/pkg/decorator"
	"github.com/sirupsen/logrus"
)

type UpdatePost struct {
	Post dto.Post
}

type UpdatePostResult struct {
	Post entity.Post
}

type UpdatePostHandler decorator.CommandHandler[UpdatePost, *UpdatePostResult]

type updatePostHandler struct {
	Repo repo.PostRepository
}

func (u updatePostHandler) Handle(ctx context.Context, command UpdatePost) (*UpdatePostResult, error) {
	p := command.Post
	if err := p.Validate(); err != nil {
		return nil, err
	}
	res, err := u.Repo.Update(ctx, p.ToEntity())
	if err != nil {
		return nil, err
	}
	return &UpdatePostResult{Post: *res}, nil
}

func NewUpdatePostHandler(
	repo repo.PostRepository,
	entry *logrus.Entry,
	record decorator.MetricsRecord,
) UpdatePostHandler {
	if repo == nil {
		entry.Panicf("post repo is nil")
	}
	return decorator.ApplyHandlerDecorators[UpdatePost, *UpdatePostResult](
		&updatePostHandler{Repo: repo},
		entry,
		record,
	)
}
