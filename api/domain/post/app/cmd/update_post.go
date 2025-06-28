package cmd

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/app/dto"
	"github.com/SimonMorphy/grog/api/domain/post/entity"
	"github.com/SimonMorphy/grog/api/domain/post/repo"
	"github.com/SimonMorphy/grog/api/infra/pkg/decorator"
)

type UpdatePost struct {
	Post dto.Post
}

type UpdatePostResult struct {
	Post entity.Post
}

type UpdatePostHandler decorator.Handler[UpdatePost, *UpdatePostResult]

type updatePostHandler struct {
	Repo repo.PostRepository
}

func (u updatePostHandler) Handle(ctx context.Context, query UpdatePost) (*UpdatePostResult, error) {
	//TODO implement me
	panic("implement me")
}
