package app

import (
	"github.com/SimonMorphy/grog/api/infra/pkg/decorator"
	"github.com/SimonMorphy/grog/api/infra/storage"
	"github.com/SimonMorphy/grog/api/post/adapters"
	"github.com/SimonMorphy/grog/api/post/app/cmd"
	"github.com/SimonMorphy/grog/api/post/app/query"
	"github.com/sirupsen/logrus"
)

type App struct {
	C Cmd
	Q Qry
}

func NewApp() App {
	DB := storage.NewPostgres()
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
