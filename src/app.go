package src

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yeisonLucio/shopping-cart/src/config"
)

type App struct {
	Routes *Routes
}

func (a *App) Start() {
	app := fiber.New()
	a.Routes.SetRoutes(app)
	app.Listen(":" + config.App.ServerPort)
}

func NewApp(routes *Routes) *App {
	return &App{
		Routes: routes,
	}
}
