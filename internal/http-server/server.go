package http_server

import (
	"backend-bootcamp-assignment-2024/internal/core"
	"backend-bootcamp-assignment-2024/internal/http-server/handler"
	"backend-bootcamp-assignment-2024/internal/model/dto/response"
	"backend-bootcamp-assignment-2024/internal/model/entity"
	"backend-bootcamp-assignment-2024/internal/pkg/web"
	"backend-bootcamp-assignment-2024/internal/service"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"slices"
	"strings"
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

	app.Router.GET("/dummyLogin", app.mappedHandler(handler.DummyLogin))
	app.Router.POST("/login", app.mappedHandler(handler.Login))
	app.Router.POST("/register", app.mappedHandler(handler.Register))

	authOnly := app.Router.Group("/")
	authOnly.Use(authMiddleware())
	{
		authOnly.GET("/house/:id", authMiddleware(), app.mappedHandler(handler.GetHouse))
		authOnly.POST("/flat/create", app.mappedHandler(handler.CreateFlat))
		authOnly.POST("/house/:id/subscribe", app.mappedHandler(handler.SubscribeHouse))
	}

	moderOnly := app.Router.Group("/")
	moderOnly.Use(authModerMiddleware())
	{
		moderOnly.POST("/house/create", app.mappedHandler(handler.CreateHouse))
		moderOnly.POST("/flat/update", app.mappedHandler(handler.UpdateFlat))
	}
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

func authModerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := getToken(ctx)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		userRole := claims["role"].(string)
		if ok && token.Valid && userRole == entity.USERTYPE_MODERATOR {
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := getToken(ctx)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		userRole := claims["role"].(string)
		if ok && token.Valid && slices.Contains(entity.UserTypes, userRole) {
			ctx.Set("User-Type", userRole)
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func getToken(ctx *gin.Context) (*jwt.Token, error) {
	bearerToken := ctx.GetHeader("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) != 2 {
		return nil, errors.New("invalid token")
	}
	tokenString := splitToken[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		//TODO: change secret
		return []byte("secret"), nil
	})
	return token, err
}
