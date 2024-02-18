package controller

import (
	"digicert/common"
	"digicert/dto"
	"digicert/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookController struct {
	bookSvc service.IBookService
}

func NewBookController(bookService service.IBookService) BookController {
	return BookController{bookSvc: bookService}
}

// Create godoc
// @Summary      Creates a book
// @Description  This API Creates a new book
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param book body dto.CreateBook true "New Book"
// @Success      200  {object}  dto.BookResponse
// @Failure      400  {object}  common.ErrorResponse
// @Failure      404  {object}  common.ErrorResponse
// @Failure      500  {object}  common.ErrorResponse
// @Router       /add [POST]
func (ctrl BookController) Create(context *gin.Context) {
	var book dto.CreateBook

	if err := context.ShouldBind(&book); err != nil {
		context.JSON(http.StatusBadRequest, common.ErrorResponse{Message: common.GetErrorMsg(err)})
		return
	}

	bookResponse, err := ctrl.bookSvc.Create(book)
	if err != nil {
		context.JSON(http.StatusBadRequest, common.ErrorResponse{Message: err.Error()})
		return
	}

	context.JSON(http.StatusOK, bookResponse)
}

// Read godoc
// @Summary      List the books
// @Description  This API List the books
// @Tags         Books
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.BookResponse
// @Failure      400  {object}  common.ErrorResponse
// @Failure      404  {object}  common.ErrorResponse
// @Failure      500  {object}  common.ErrorResponse
// @Router       /read [GET]
func (ctrl BookController) Read(context *gin.Context) {

	bookResponse, err := ctrl.bookSvc.Read()
	if err != nil {
		if len(bookResponse) == 0 {
			context.JSON(http.StatusNotFound, common.ErrorResponse{Message: "Not Found"})
			return
		}

		context.JSON(http.StatusBadRequest, common.ErrorResponse{Message: err.Error()})
		return
	}

	context.JSON(http.StatusOK, bookResponse)
}

// Update godoc
// @Summary      Updates a book
// @Description  This API Updates a book
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        id path string true "book id"
// @Param book body dto.UpdateBook true "Existing Book"
// @Success      200  {object} common.ErrorResponse
// @Failure      400  {object}  common.ErrorResponse
// @Failure      404  {object}  common.ErrorResponse
// @Failure      500  {object}  common.ErrorResponse
// @Router       /update/{id} [PUT]
func (ctrl BookController) Update(context *gin.Context) {

	bookId := context.Param("id")
	bookRequest := dto.UpdateBook{}

	if err := context.ShouldBind(&bookRequest); err != nil {
		context.JSON(http.StatusBadRequest, common.ErrorResponse{Message: common.GetErrorMsg(err)})
		return
	}

	err := ctrl.bookSvc.Update(bookId, bookRequest)
	if err != nil {
		if err.Error() == "not found" {
			context.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}
	context.JSON(http.StatusOK, common.ErrorResponse{Message: "Success"})

}

// Delete godoc
// @Summary      Deletes a book
// @Description  This API Deletes book
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        id path string true "book id"
// @Success      200  {object}  dto.BookResponse
// @Failure      400  {object}  common.ErrorResponse
// @Failure      404  {object}  common.ErrorResponse
// @Failure      500  {object}  common.ErrorResponse
// @Router       /delete/{id} [DELETE]
func (ctrl BookController) Delete(context *gin.Context) {
	id := context.Param("id")
	err := ctrl.bookSvc.Delete(id)

	if err != nil {
		if err.Error() == "not found" {
			context.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	context.JSON(http.StatusOK, common.ErrorResponse{Message: "Success"})
}

// GetById godoc
// @Summary      Gets a book
// @Description  This API Gets a book by ID
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        id path string true "book id"
// @Success      200  {object}  dto.BookResponse
// @Failure      400  {object}  common.ErrorResponse
// @Failure      404  {object}  common.ErrorResponse
// @Failure      500  {object}  common.ErrorResponse
// @Router       /getbyid/{id} [GET]
func (ctrl BookController) GetById(context *gin.Context) {
	id := context.Param("id")
	bookResponse, err := ctrl.bookSvc.GetById(id)
	if err != nil {
		if err.Error() == "not found" {
			context.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	context.JSON(http.StatusOK, bookResponse)
}
