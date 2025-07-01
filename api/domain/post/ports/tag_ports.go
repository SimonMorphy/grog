package ports

import "github.com/gin-gonic/gin"

type TagPorts interface {
	GetTag(ctx *gin.Context)
	ListTag(ctx *gin.Context)
	CreateTag(ctx *gin.Context)
	UpdateTag(ctx *gin.Context)
	DeleteTag(ctx *gin.Context)
}
