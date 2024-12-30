// @title           Go Backend Template
// @version         1.0
// @description     My Go Backend Template
// @host            localhost:8080
// @BasePath        /
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name x-token
package main

import (
	"log"
	"net/http"

	"github.com/eiei114/go-backend-template/application/middleware"
	"github.com/eiei114/go-backend-template/application/service"
	"github.com/eiei114/go-backend-template/config" // ここでパッケージをインポート
	_ "github.com/eiei114/go-backend-template/docs"
	infrastructure "github.com/eiei114/go-backend-template/infrastructure/persistence"
	"github.com/eiei114/go-backend-template/interface/handler"
	"github.com/eiei114/go-backend-template/interface/router"
)

func main() {
	db, _ := config.NewDBConnection()
	//database_init.CreateTable(db)

	userRepository := infrastructure.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(*userService)
	middleware := middleware.NewMiddleware(*userService)
	router := router.NewRouter(*userHandler, *middleware)
	r := router.InitRouter()

	log.Println("listening on http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", r))
}
