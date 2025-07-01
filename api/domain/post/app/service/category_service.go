package service

import (
	"github.com/SimonMorphy/grog/api/domain/post/app"
	"github.com/SimonMorphy/grog/api/domain/post/app/cmd"
	"github.com/SimonMorphy/grog/api/domain/post/entity"
	"github.com/SimonMorphy/grog/api/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CategoryService struct {
	APP  app.App
	RESP types.Response
}

func (c CategoryService) CreateCategory(ctx *gin.Context) {
	var categories []*entity.Category
	err := ctx.ShouldBindJSON(&categories)
	if err != nil {
		logrus.Error("Error creating category", err)
		c.RESP.Error(ctx, err)
		return
	}
	c.APP.C.CreateCategory.Handle(ctx.Request.Context(), cmd.CreateCategory{})

}

func (c CategoryService) UpdateCategory(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) DeleteCategory(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) ListCategory(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryService) GetCategory(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
