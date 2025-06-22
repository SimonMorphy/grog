package main

import (
	_ "github.com/SimonMorphy/grog/api/infra/config"
	"github.com/SimonMorphy/grog/api/post"
	"github.com/SimonMorphy/grog/api/post/app"
	"github.com/SimonMorphy/grog/api/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	engine := gin.New()
	resp := types.Response{}
	App := app.NewApp()
	server := post.HttpServer{
		Resp: resp,
		App:  App,
	}
	server.Apply(engine)
	err := engine.Run(":8080")
	if err != nil {
		logrus.Panicf(
			"Failed to start server: %v", err)
	}
}
