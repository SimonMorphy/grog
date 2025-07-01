package adapters

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/entity"
	"gorm.io/gorm"
)

type RepositoryCategory struct {
	DB *gorm.DB
}

func (r RepositoryCategory) BatchSave(ctx context.Context, categories []*entity.Category) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		return r.DB.WithContext(ctx).Create(categories).Error
	})
}

func (r RepositoryCategory) Create(ctx context.Context, e *entity.Category) (*entity.Category, error) {
	if err := r.DB.Model(CATEGORY).WithContext(ctx).Create(e).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func (r RepositoryCategory) Get(ctx context.Context, u uint) (c *entity.Category, err error) {
	err = r.DB.Model(CATEGORY).WithContext(ctx).First(c, u).Error
	return
}

func (r RepositoryCategory) List(ctx context.Context) (categories []*entity.Category, err error) {
	err = r.DB.Model(CATEGORY).WithContext(ctx).Find(&categories).Error
	return
}

func (r RepositoryCategory) Update(ctx context.Context, e *entity.Category) (*entity.Category, error) {
	err := r.DB.Model(CATEGORY).WithContext(ctx).Save(e).Error
	return e, err
}

func (r RepositoryCategory) Delete(ctx context.Context, u uint) error {
	return r.DB.WithContext(ctx).Delete(CATEGORY, u).Error
}
