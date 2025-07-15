package dto

type Page struct {
	Page int `form:"page" json:"page" validate:"required,min=1"`
	Size int `form:"size" json:"size" validate:"required,min=1,max=50"`
}

func (p *Page) Validate() error {
	return v.Struct(p)
}
