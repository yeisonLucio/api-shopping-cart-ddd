package src

import (
	"github.com/gin-gonic/gin"
	"github.com/yeisonLucio/shopping-cart/src/config"
)

type App struct {
	Routes *Routes
}

func (a *App) Start() {
	app := gin.New()
	a.Routes.SetRoutes(app)
	app.Run(":" + config.App.ServerPort)
	app.Use(gin.Logger())
}

func NewApp(routes *Routes) *App {
	return &App{
		Routes: routes,
	}
}
