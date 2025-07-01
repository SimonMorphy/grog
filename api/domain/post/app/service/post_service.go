package service

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

type PostService struct {
	APP  app.App
	RESP types.Response
}

func (p PostService) CreatePost(ctx *gin.Context) {
	var post dto.Post
	err := ctx.ShouldBind(&post)
	if err != nil {
		logrus.Error("Error binding post", err)
		p.RESP.Error(ctx, err)
		return
	}
	err = post.Validate()
	if err != nil {
		logrus.Error("Error invalid post", err)
		p.RESP.Error(ctx, err)
		return
	}
	result, err := p.APP.C.CreatePost.Handle(ctx.Request.Context(), cmd.CreatePost{Post: post})
	if err != nil {
		logrus.Error("Error creating post", err)
		p.RESP.Error(ctx, err)
		return
	}
	p.RESP.Success(ctx, result)
	return
}

func (p PostService) GetPost(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		logrus.Error("Error parsing id", err)
		p.RESP.Error(ctx, err)
		return
	}
	post, err := p.APP.Q.GetPost.Handle(ctx.Request.Context(), query.GetPost{ID: id})
	if err != nil {
		logrus.Error("Error getting post", err)
		p.RESP.Error(ctx, err)
		return
	}
	p.RESP.Success(ctx, post)
	return
}

func (p PostService) UpdatePost(ctx *gin.Context) {
	var post dto.Post
	err := ctx.ShouldBind(&post)
	if err != nil {
		logrus.Error("Error binding post", err)
		p.RESP.Error(ctx, err)
		return
	}
	result, err := p.APP.C.UpdatePost.Handle(ctx.Request.Context(), cmd.UpdatePost{Post: post})
	if err != nil {
		logrus.Error("Error updating post", err)
		p.RESP.Error(ctx, err)
		return
	}
	p.RESP.Success(ctx, result)
	return
}

func (p PostService) DeletePost(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		logrus.Error("Error parsing id", err)
		p.RESP.Error(ctx, err)
		return
	}
	result, err := p.APP.C.DeletePost.Handle(ctx.Request.Context(), cmd.DeletePost{ID: id})
	if err != nil {
		logrus.Error("Error deleting post", err)
		p.RESP.Error(ctx, err)
		return
	}
	p.RESP.Success(ctx, result)
	return
}

func (p PostService) ListPosts(ctx *gin.Context) {
	posts, err := p.APP.Q.ListPost.Handle(ctx.Request.Context(), query.ListPost{})
	if err != nil {
		logrus.Error("Error listing posts", err)
		p.RESP.Error(ctx, err)
		return
	}
	p.RESP.Success(ctx, posts)
	return
}
