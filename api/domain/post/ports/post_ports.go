package ports

import (
	"github.com/gin-gonic/gin"
)

type PostPort interface {
	CreatePost(ctx *gin.Context)
	GetPost(ctx *gin.Context)
	UpdatePost(ctx *gin.Context)
	DeletePost(ctx *gin.Context)
	ListPosts(ctx *gin.Context)
}
