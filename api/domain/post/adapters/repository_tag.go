package adapters

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/entity"
	"gorm.io/gorm"
)

type RepositoryTag struct {
	DB *gorm.DB
}

func (r RepositoryTag) Create(ctx context.Context, e *entity.Tag) (*entity.Tag, error) {
	if err := r.DB.WithContext(ctx).Create(e).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func (r RepositoryTag) Get(ctx context.Context, u uint) (t *entity.Tag, err error) {
	err = r.DB.WithContext(ctx).First(t, u).Error
	return
}

func (r RepositoryTag) List(ctx context.Context) (ts []*entity.Tag, err error) {
	err = r.DB.WithContext(ctx).Model(TAG).Preload("Posts").Find(&ts).Error
	return
}

func (r RepositoryTag) Update(ctx context.Context, e *entity.Tag) (*entity.Tag, error) {
	err := r.DB.WithContext(ctx).Model(TAG).Updates(e).Error
	return e, err
}

func (r RepositoryTag) Delete(ctx context.Context, u uint) error {
	return r.DB.WithContext(ctx).Model(TAG).Delete(u).Error
}

func (r RepositoryTag) BatchSave(ctx context.Context, tags []*entity.Tag) error {
	//TODO implement me
	panic("implement me")
}
