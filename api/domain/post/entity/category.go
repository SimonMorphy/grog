package entity

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string  `json:"name" gorm:"size:50;not null;unique"`
	Posts []*Post `json:"posts" gorm:"many2many:post_categories;"`
}

func (c Category) TableName() string {
	return "category"
}
