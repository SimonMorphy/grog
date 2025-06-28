package ports

import (
	"github.com/gin-gonic/gin"
)

type PostPort interface {
	CreatePost(ctx *gin.Context)
	GetPost(ctx *gin.Context)
}
