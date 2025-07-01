package repo

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/entity"
)

type TagRepository interface {
	CommonRepository[entity.Tag]
	BatchSave(context.Context, []*entity.Tag) error
}
