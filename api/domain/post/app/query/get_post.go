package query

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/entity"
	"github.com/SimonMorphy/grog/api/domain/post/repo"
	"github.com/SimonMorphy/grog/api/infra/pkg/decorator"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GetPost struct {
	ID uint
}

type GetPostResponse struct {
	Post *entity.Post
}

type GetPostHandler decorator.Handler[GetPost, *GetPostResponse]
type getPostHandler struct {
	Repository repo.PostRepository
}

func (g getPostHandler) Handle(ctx context.Context, query GetPost) (*GetPostResponse, error) {
	post, err := g.Repository.Get(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	return &GetPostResponse{
		Post: post,
	}, nil
}
func NewGetPostHandler(
	repository repo.PostRepository,
	entry *logrus.Entry,
	record decorator.MetricsRecord,
) GetPostHandler {
	if repository == nil {
		logrus.Panic(gorm.ErrInvalidDB)
	}
	return decorator.ApplyHandlerDecorators[GetPost, *GetPostResponse](
		&getPostHandler{Repository: repository},
		entry,
		record,
	)

}
