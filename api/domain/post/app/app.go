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
	Repo := adapters.NewRepositoryPostgres(DB)
	logger := logrus.NewEntry(logrus.StandardLogger())
	todoMetrics := decorator.NewToDoMetrics()
	return App{
		C: Cmd{
			CreatePost: cmd.NewCreatePostHandler(Repo, logger, todoMetrics),
		},
		Q: Qry{
			GetPost: query.NewGetPostHandler(Repo, logger, todoMetrics),
		},
	}
}

type Cmd struct {
	CreatePost cmd.CreatePostHandler
}

type Qry struct {
	GetPost query.GetPostHandler
}
