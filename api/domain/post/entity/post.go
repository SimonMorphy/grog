package entity

import (
	"fmt"
	"github.com/SimonMorphy/grog/api/domain/post/vo"
	"golang.org/x/crypto/bcrypt"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title           string         `json:"title" gorm:"size:255;not null"`
	Content         string         `json:"content" gorm:"type:text;not null"`
	Description     string         `json:"description" gorm:"size:1023"`
	LeadImage       string         `json:"lead_image" gorm:"size:1023"`
	Status          vo.PostStatus  `json:"status" gorm:"size:31;not null;default:0"`
	IsRecommend     bool           `json:"is_recommend" gorm:"not null;default:false"`
	IsCommentEnable bool           `json:"is_comment_enable" gorm:"not null;default:true"`
	IsTop           bool           `json:"is_top" gorm:"not null;default:false"`
	Views           uint           `json:"views" gorm:"not null;default:0"`
	Word            uint           `json:"word" gorm:"not null;default:0"`
	PublishTime     *time.Time     `json:"publish_time"`
	ReadTime        *time.Duration `json:"read_time"`
	Password        string         `json:"-" gorm:"size:255"`
	Categories      []*Category    `json:"categories" gorm:"many2many:post_categories;"`
	Tags            []*Tag         `json:"tags" gorm:"many2many:post_tags;"`
}

func (p *Post) TableName() string {
	return "post"
}

func (p *Post) EncryptPassword() error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	p.Password = string(hashedBytes)
	return err
}

func (p *Post) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
}

func (p *Post) TransitionTo(newStatus vo.PostStatus) error {
	currentStatus := p.Status

	if allowed, exists := vo.StatusTransitions[currentStatus][newStatus]; !exists || !allowed {
		return fmt.Errorf("invalid cast :-> %s → %s", currentStatus, newStatus)
	}

	switch {
	case currentStatus == vo.Scheduled && newStatus == vo.Published:
		*p.PublishTime = time.Now()
	case newStatus == vo.Deleted:
		p.Model.DeletedAt = gorm.DeletedAt{Time: time.Now()}
	}

	p.Status = newStatus
	return nil
}
