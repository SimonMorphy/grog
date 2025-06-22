package repo

import (
	"context"
	"github.com/SimonMorphy/grog/api/post/entity"
)

type PostRepository interface {
	Create(context.Context, *entity.Post) (*entity.Post, error)
	Get(context.Context, uint) (*entity.Post, error)
}
