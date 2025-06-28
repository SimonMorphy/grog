package repo

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/app/dto"
	"github.com/SimonMorphy/grog/api/domain/post/entity"
)

type PostRepository interface {
	Create(context.Context, *entity.Post) (*entity.Post, error)
	Get(context.Context, uint) (*entity.Post, error)
	List(context.Context) ([]*entity.Post, error)
	Update(context.Context, *dto.Post) (*entity.Post, error)
	Delete(context.Context, uint) error
}

type CategoryRepository interface {
	Create(context.Context, *entity.Category) (*entity.Category, error)
	Get(context.Context, uint) (*entity.Category, error)
	List(context.Context) ([]*entity.Category, error)
	Update(context.Context, *entity.Category) (*entity.Category, error)
	Delete(context.Context, uint) error
}
