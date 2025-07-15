package repo

import (
	"context"
)

type CommonRepository[E any] interface {
	Create(context.Context, *E) (*E, error)
	Get(context.Context, uint) (*E, error)
	List(context.Context, int, int) ([]*E, error)
	Update(context.Context, *E) (*E, error)
	Delete(context.Context, uint) error
}
