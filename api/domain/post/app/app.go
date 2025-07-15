package app

import (
	"github.com/SimonMorphy/grog/api/domain/post/adapters"
	"github.com/SimonMorphy/grog/api/domain/post/app/cmd"
	"github.com/SimonMorphy/grog/api/domain/post/app/query"
	"github.com/SimonMorphy/grog/api/domain/post/entity"
	"github.com/SimonMorphy/grog/api/infra/pkg/decorator"
	"github.com/SimonMorphy/grog/api/infra/storage"
	"github.com/sirupsen/logrus"
)

type App struct {
	C Cmd
	Q Qry
}

func NewApp() App {
	DB := storage.NewPostgres()
	err := DB.AutoMigrate(&entity.Post{}, &entity.Tag{}, &entity.Category{})
	if err != nil {
		logrus.Fatalf("failed to migrate database: %v", err)
	}
	postRepo := adapters.NewRepositoryPost(DB)
	cateRepo := adapters.NewRepositoryCategory(DB)
	logger := logrus.NewEntry(logrus.StandardLogger())
	todoMetrics := decorator.NewToDoMetrics()
	return App{
		C: Cmd{
			CreatePost:     cmd.NewCreatePostHandler(postRepo, logger, todoMetrics),
			CreateCategory: cmd.NewCreateCategoryHandler(cateRepo, logger, todoMetrics),
			DeletePost:     cmd.NewDeletePostHandler(postRepo, logger, todoMetrics),
			UpdatePost:     cmd.NewUpdatePostHandler(postRepo, logger, todoMetrics),
		},
		Q: Qry{
			GetPost:  query.NewGetPostHandler(postRepo, logger, todoMetrics),
			ListPost: query.NewListPostHandler(postRepo, logger, todoMetrics),
		},
	}
}

type Cmd struct {
	CreatePost     cmd.CreatePostHandler
	CreateCategory cmd.CreateCategoryHandler
	DeleteCategory cmd.DeleteCategoryHandler
	DeletePost     cmd.DeletePostHandler
	UpdatePost     cmd.UpdatePostHandler
}

type Qry struct {
	GetPost     query.GetPostHandler
	ListPost    query.ListPostHandler
	GetCategory query.GetCategoryHandler
}
