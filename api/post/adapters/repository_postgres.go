package adapters

import (
	"context"
	"github.com/SimonMorphy/grog/api/post/entity"
	"gorm.io/gorm"
)

type RepositoryPostgres struct {
	DB *gorm.DB
}

func NewRepositoryPostgres(DB *gorm.DB) *RepositoryPostgres {
	return &RepositoryPostgres{DB: DB}
}

func (r RepositoryPostgres) Create(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	r.DB.WithContext(ctx).Create(post)
	if post.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return post, nil
}

func (r RepositoryPostgres) Get(ctx context.Context, id uint) (*entity.Post, error) {
	var post entity.Post
	if err := r.DB.WithContext(ctx).First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
