package repo

import (
	"github.com/SimonMorphy/grog/api/domain/post/entity"
)

type PostRepository interface {
	CommonRepository[entity.Post]
}
