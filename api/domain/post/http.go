package post

import (
	"github.com/SimonMorphy/grog/api/domain/post/app"
	"github.com/SimonMorphy/grog/api/domain/post/app/cmd"
	"github.com/SimonMorphy/grog/api/domain/post/app/dto"
	"github.com/SimonMorphy/grog/api/domain/post/app/query"
	"github.com/SimonMorphy/grog/api/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

// HttpServer handles HTTP requests related to posts
type HttpServer struct {
	Resp types.Response
	App  app.App
}

func (h HttpServer) CreatePost(ctx *gin.Context) {
	var post dto.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		logrus.Errorf("Error binding post -> %v", err)
		h.Resp.Error(ctx, err)
		return
	}
	handle, err := h.App.C.CreatePost.Handle(ctx.Request.Context(), cmd.CreatePost{Post: post})
	if err != nil {
		logrus.Errorf("Error create post -> %v", err)
		h.Resp.Error(ctx, err)
	}
	h.Resp.Success(ctx, handle)
}

func (h HttpServer) GetPost(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		logrus.Errorf("Error parsing id -> %s", err.Error())
		h.Resp.Error(ctx, err)
		return
	}
	rsp, err := h.App.Q.GetPost.Handle(ctx.Request.Context(), query.GetPost{ID: uint(id)})
	if err != nil {
		logrus.Errorf("Error getting post by id -> %s", err.Error())
		h.Resp.Error(ctx, err)
		return
	}
	h.Resp.Success(ctx, rsp)
}

func (h HttpServer) UpdatePost(ctx *gin.Context) {
	var post dto.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		logrus.Errorf("Error parsing post: %s", err.Error())
		h.Resp.Error(ctx, err)
		return
	}
	result, err := h.App.C.UpdatePost.Handle(ctx.Request.Context(), cmd.UpdatePost{Post: post})
	if err != nil {
		logrus.Errorf("Error updating post: %s", err.Error())
		h.Resp.Error(ctx, err)
		return
	}
	h.Resp.Success(ctx, result)
}

func (h HttpServer) DeletePost(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		logrus.Errorf("Error parsing id -> %s", err.Error())
		h.Resp.Error(ctx, err)
		return
	}
	result, err := h.App.C.DeletePost.Handle(ctx.Request.Context(), cmd.DeletePost{ID: uint(id)})
	if err != nil {
		logrus.Errorf("Error deleting post by id -> %s", err.Error())
		h.Resp.Error(ctx, err)
		return
	}
	h.Resp.Success(ctx, result)
}

func (h HttpServer) ListPosts(ctx *gin.Context) {
	var page dto.Page
	if err := ctx.ShouldBindQuery(&page); err != nil {
		logrus.Errorf("invalid page query, use ?page=&size= to list query, err -> %v", err)
		h.Resp.Error(ctx, err)
		return
	}
	if err := page.Validate(); err != nil {
		logrus.Errorf("invalid page param,err -> %v", err)
		h.Resp.Error(ctx, err)
		return
	}
	result, err := h.App.Q.ListPost.Handle(ctx.Request.Context(), query.ListPost{Page: page.Page, Size: page.Size})
	if err != nil {
		logrus.Errorf("Error listing post -> %s", err.Error())
		h.Resp.Error(ctx, err)
		return
	}
	h.Resp.Success(ctx, result)
}

// Apply registers Post routes to the router
func (h HttpServer) Apply(e *gin.Engine) {
	postGroup := e.Group("/post")
	{
		postGroup.POST("/", h.CreatePost)
		postGroup.GET("/:id", h.GetPost)
		postGroup.PUT("/", h.UpdatePost)
		postGroup.DELETE("/:id", h.DeletePost)
		postGroup.GET("/list", h.ListPosts)
	}
}
