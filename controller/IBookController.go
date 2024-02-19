package controller

import "github.com/gin-gonic/gin"

type IBookController interface {
	Create(context *gin.Context)
	Read(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	GetById(context *gin.Context)
}
