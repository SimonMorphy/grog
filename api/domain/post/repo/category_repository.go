package repo

import (
	"context"
	"github.com/SimonMorphy/grog/api/domain/post/entity"
)

type CategoryRepository interface {
	CommonRepository[entity.Category]
	BatchSave(context.Context, []*entity.Category) error
}
