package router

import (
	"github.com/eiei114/go-backend-template/application/middleware"
	"github.com/eiei114/go-backend-template/interface/handler"
	"github.com/uptrace/bunrouter"
)

type Router struct {
	UserHandler handler.UserHandler
	Middleware  middleware.Middleware
}

func NewRouter(userHandler handler.UserHandler, middleware middleware.Middleware) *Router {
	return &Router{
		UserHandler: userHandler,
		Middleware:  middleware,
	}
}

func (i *Router) InitRouter() *bunrouter.Router {
	b := bunrouter.New()
	b.Use(i.Middleware.RecoverMiddleware())
	b.Use(i.Middleware.CorsMiddleware())

	b.POST("/user/create", i.UserHandler.UserCreateHandle())

	b.Use(i.Middleware.AuthenticateMiddleware()).WithGroup("", func(group *bunrouter.Group) {
		group.POST("/user/get", i.UserHandler.UserGetHandle())
		group.POST("/user/count", i.UserHandler.CountAddHandle())
		group.POST("/user/destroy", i.UserHandler.DestroyHandle())
	})

	return b
}