package main

import (
	"context"
	"digicert/config"
	"digicert/controller"
	"digicert/db"
	"digicert/repository"
	"digicert/service"
)

var (
	ctx  context.Context
	cnf  = config.LoadConfig()
	svc  service.IBookService
	repo repository.IBookRepository
	ctrl controller.BookController
)

func init() {
	db, _ := db.Connect(ctx, cnf)
	repo = repository.NewBookRepository(ctx, db)
	svc = service.NewBookService(repo)
	ctrl = controller.NewBookController(svc)
}
