package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ramadhanalfarisi/go-crawler/controller"
)

type App struct {
	Routes *gin.Engine
}

func (a *App) CreateRoutes() {
	g := gin.Default()
	g.GET("/manga", controller.GetDataManga)
	a.Routes = g
}

func (a *App) Run(){
	a.Routes.Run(":8080")
}