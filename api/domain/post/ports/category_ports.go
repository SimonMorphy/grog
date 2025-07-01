package ports

import "github.com/gin-gonic/gin"

type CategoryPort interface {
	CreateCategory(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
	ListCategory(ctx *gin.Context)
	GetCategory(ctx *gin.Context)
}
