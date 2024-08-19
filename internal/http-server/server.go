package http_server

import (
	"backend-bootcamp-assignment-2024/internal/core"
	"backend-bootcamp-assignment-2024/internal/http-server/handler"
	"backend-bootcamp-assignment-2024/internal/model/dto/response"
	"backend-bootcamp-assignment-2024/internal/pkg/web"
	"backend-bootcamp-assignment-2024/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type App struct {
	Server  web.Server
	Router  *gin.Engine
	Service *service.Service
}

func New(service *service.Service, cfg *core.Config) *App {
	app := &App{
		Service: service,
	}
	app.initRoutes()
	app.Server = web.NewServer(cfg.Server, app.Router)
	return app
}

func (app *App) Start(ctx context.Context) error {
	return app.Server.Run(ctx)
}

func (app *App) initRoutes() {
	app.Router = gin.Default()

	//app.Router.GET("/dummyLogin", app.mappedHandler())
	//app.Router.POST("/login", app.mappedHandler())
	//app.Router.POST("/register", app.mappedHandler())
	app.Router.GET("/house/:id", app.mappedHandler(handler.GetHouse))
	app.Router.POST("/house/:id/subscribe", app.mappedHandler(handler.SubscribeHouse))
	app.Router.POST("/flat/create", app.mappedHandler(handler.CreateFlat))
	app.Router.POST("/house/create", app.mappedHandler(handler.CreateHouse))
	app.Router.POST("/flat/update", app.mappedHandler(handler.UpdateFlat))
}

func (app *App) mappedHandler(handler func(*gin.Context, *service.Service) error) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		if err := handler(ctx, app.Service); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.Error{
				Code:      http.StatusInternalServerError,
				Message:   err.Error(),
				RequestId: ctx.GetHeader("X-Request-Id"),
			})
			ctx.Set("Retry-After", 10)
			ctx.Abort()
		}
	}
}
