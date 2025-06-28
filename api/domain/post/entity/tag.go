package entity

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name  string  `gorm:"size:50;not null"`
	Posts []*Post `gorm:"many2many:post_tags;"`
}

func (t Tag) TableName() string {
	return "tag"
}
