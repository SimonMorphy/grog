package entity

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string  `gorm:"size:50;not null"`
	Posts []*Post `gorm:"many2many:post_categories;"`
}

func (c Category) TableName() string {
	return "category"
}
