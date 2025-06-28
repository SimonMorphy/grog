package adapters

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/entity"
	"gorm.io/gorm"
)


type RepositoryCategory struct {
	DB *gorm.DB
}

func (r RepositoryCategory) Create(ctx context.Context, e *entity.Category) (*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryCategory) Get(ctx context.Context, u uint) (*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryCategory) List(ctx context.Context) ([]*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryCategory) Update(ctx context.Context, e *entity.Category) (*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryCategory) Delete(ctx context.Context, u uint) error {
	//TODO implement me
	panic("implement me")
}
