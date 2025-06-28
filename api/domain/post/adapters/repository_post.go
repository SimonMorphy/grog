package adapters

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/app/dto"
	"github.com/SimonMorphy/grog/api/domain/post/entity"
	"gorm.io/gorm"
)

type RepositoryPost struct {
	DB *gorm.DB
}

func (r RepositoryPost) List(ctx context.Context) ([]*entity.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryPost) Update(ctx context.Context, post *dto.Post) (*entity.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryPost) Delete(ctx context.Context, u uint) error {
	//TODO implement me
	panic("implement me")
}

func NewRepositoryPostgres(DB *gorm.DB) *RepositoryPost {
	return &RepositoryPost{DB: DB}
}

func (r RepositoryPost) Create(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	return post, r.DB.Model(&POST).WithContext(ctx).Create(post).Error
}

func (r RepositoryPost) Get(ctx context.Context, id uint) (*entity.Post, error) {
	var post entity.Post
	if err := r.DB.WithContext(ctx).First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
