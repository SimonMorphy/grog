package dto

import (
	"github.com/SimonMorphy/grog/api/domain/post/entity"
)

// Post is a versatile data transfer object for creating and updating posts.
// Validation is kept flexible for partial updates; the application service
// is responsible for enforcing stricter rules, such as for creation.
type Post struct {
	Title           string   `json:"title,omitempty" validate:"omitempty,min=3,max=255"`
	Content         string   `json:"content,omitempty" validate:"omitempty,min=10"`
	Description     string   `json:"description,omitempty" validate:"omitempty,max=1023"`
	LeadImage       string   `json:"lead_image,omitempty" validate:"omitempty,url"`
	Status          *int     `json:"status,omitempty" validate:"omitempty,oneof=0 1 2 3 4 5 6"`
	Password        string   `json:"password,omitempty" validate:"omitempty,min=6"`
	IsRecommend     *bool    `json:"is_recommend,omitempty"`
	IsCommentEnable *bool    `json:"is_comment_enable,omitempty"`
	IsTop           *bool    `json:"is_top,omitempty"`
	Categories      []string `json:"categories,omitempty" validate:"omitempty,dive,min=1"`
	Tags            []string `json:"tags,omitempty" validate:"omitempty,dive,min=1"`
}

func (p *Post) Validate() error {
	return v.Struct(p)
}

func ptrToBool(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}

// ToEntity converts Post to entity.Post, handling optional fields.
func (p *Post) ToEntity() *entity.Post {
	categories := make([]*entity.Category, 0)
	tags := make([]*entity.Tag, 0)
	for _, cateName := range p.Categories {
		categories = append(categories, &entity.Category{Name: cateName})
	}
	for _, tagName := range p.Tags {
		tags = append(tags, &entity.Tag{Name: tagName})
	}
	post := entity.Post{
		Title:           p.Title,
		Content:         p.Content,
		Description:     p.Description,
		LeadImage:       p.LeadImage,
		Status:          entity.PrtToStatus(p.Status),
		IsRecommend:     ptrToBool(p.IsRecommend),
		IsCommentEnable: ptrToBool(p.IsCommentEnable),
		IsTop:           ptrToBool(p.IsTop),
		Password:        p.Password,
		Categories:      categories,
		Tags:            tags,
	}
	return &post
}
