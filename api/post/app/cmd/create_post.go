package cmd

import (
	"context"
	"github.com/SimonMorphy/grog/api/infra/pkg/decorator"
	"github.com/SimonMorphy/grog/api/post/app/dto"
	"github.com/SimonMorphy/grog/api/post/repo"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CreatePost struct {
	Post dto.PostDTO
}

type CreatePostResult struct {
	ID uint
}

type CreatePostHandler decorator.CommandHandler[CreatePost, *CreatePostResult]

type createPostHandler struct {
	Repository repo.PostRepository
}

func (c createPostHandler) Handle(ctx context.Context, query CreatePost) (*CreatePostResult, error) {
	var err error
	post := query.Post
	if err = post.Validate(); err != nil {
		logrus.Errorf("CreatePostHandler.Handle: post validation failed: %v", err)
		return nil, err
	}
	usr, err := c.Repository.Create(ctx, post.ToEntity())
	if err != nil {
		logrus.Errorf("CreatePostHandler.Handle: failed to create post: %v", err)
		return nil, err
	}
	return &CreatePostResult{
		ID: usr.ID,
	}, nil
}

func NewCreatePostHandler(
	repo repo.PostRepository,
	entry *logrus.Entry,
	record decorator.MetricsRecord,
) CreatePostHandler {

	if repo == nil {
		logrus.Panic(gorm.ErrInvalidDB)
	}
	return decorator.ApplyHandlerDecorators[CreatePost, *CreatePostResult](
		&createPostHandler{Repository: repo},
		entry,
		record,
	)
}
