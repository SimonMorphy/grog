package dto

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

var (
	v *validator.Validate
	o = sync.Once{}
)

func init() {
	o.Do(func() {
		v = validator.New()
	})
}
