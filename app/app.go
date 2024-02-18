package app

import (
	"context"
	"digicert/config"
	"digicert/controller"
	"digicert/db"
	"digicert/docs"
	"digicert/repository"
	"digicert/routes"
	"digicert/service"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	server *gin.Engine
	ctx    context.Context
	cnf    = config.LoadConfig()
	svc    service.IBookService
	repo   repository.IBookRepository
	ctrl   controller.BookController
)

func init() {
	db, _ := db.Connect(ctx, cnf)
	repo = repository.NewBookRepository(ctx, db)
	svc = service.NewBookService(repo)
	ctrl = controller.NewBookController(svc)

	docs.SwaggerInfo.Title = "Swagger Book API"
	docs.SwaggerInfo.Description = "This is a Book API server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cnf.ServerName, cnf.ServerPort)
	docs.SwaggerInfo.BasePath = "/api/v1/books"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	fmt.Println(fmt.Sprintf("Swagger is serving on http://%s:%s/swagger/index.html", cnf.ServerName, cnf.ServerPort))
	gin.SetMode(gin.DebugMode)
	server = gin.Default()
}

func StartApplication() {
	routes.RegisterRoutes(server, ctrl)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Run(fmt.Sprintf(":%s", cnf.ServerPort))
}
