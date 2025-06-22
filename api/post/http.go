package post

import (
	"github.com/SimonMorphy/grog/api/post/app"
	"github.com/SimonMorphy/grog/api/post/app/cmd"
	"github.com/SimonMorphy/grog/api/post/app/dto"
	"github.com/SimonMorphy/grog/api/types"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Resp types.Response
	App  app.App
}

func (h HttpServer) CreatePost(ctx *gin.Context) {
	var post dto.PostDTO
	if err := ctx.ShouldBindJSON(&post); err != nil {
		h.Resp.Error(ctx, err)
	}
	handle, err := h.App.C.CreatePost.Handle(ctx.Request.Context(), cmd.CreatePost{Post: post})
	if err != nil {
		h.Resp.Error(ctx, err)
	}
	h.Resp.Success(ctx, handle)
}

func (h HttpServer) GetPost(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h HttpServer) Apply(e *gin.Engine) {
	postGroup := e.Group("/post")
	{
		postGroup.POST("/", h.CreatePost)
		postGroup.GET("/:id", h.GetPost)
	}

}
