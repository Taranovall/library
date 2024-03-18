package book_handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	book_controllers "library-service/controllers/book-controllers"
	"library-service/jwt"
	"library-service/utils"
	"net/http"
	"strconv"
	"time"
)

type bookHandler struct {
	service book_controllers.Service
}

func NewBookHandler(service book_controllers.Service) *bookHandler {
	return &bookHandler{service: service}
}

func (h *bookHandler) CreateHandler(context *gin.Context) {

	var book book_controllers.BookInput

	context.Bind(&book)

	result, statusCode, err := h.service.CreateBook(&book)

	switch statusCode {
	case http.StatusCreated:
		utils.APIResponse(context, "Book created successfully.", http.StatusCreated, http.MethodPost, result)
		return

	case http.StatusInternalServerError:
		utils.APIResponse(context, "Internal Server error occured", http.StatusInternalServerError, http.MethodPost, err)
		return
	}

}

func (h *bookHandler) DeleteHandler(context *gin.Context) {

	idStr := context.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		logrus.Warnf("Malformed parameter. %s", err.Error())
		utils.ValidatorErrorResponse(context, http.StatusBadRequest, http.MethodDelete, err.Error())
		return
	}

	statusCode, err := h.service.DeleteBook(uint(id))

	switch statusCode {
	case http.StatusNoContent:
		utils.APIResponse(context, "Book deleted successfully.", http.StatusCreated, http.MethodDelete, nil)
		return

	case http.StatusNotFound:
		utils.APIResponse(context, "Book with specified ID wasn't found", http.StatusNotFound, http.MethodDelete, err)
		return

	case http.StatusInternalServerError:
		utils.APIResponse(context, "Internal server error occurred", http.StatusInternalServerError, http.MethodDelete, err)
	}
}

func (h *bookHandler) GetAllHandler(context *gin.Context) {
	result, statusCode, err := h.service.FindAllBooks()

	switch statusCode {
	case http.StatusOK:
		utils.APIResponse(context, "Book has been retrieved successfully.", http.StatusOK, http.MethodGet, result)
		return

	case http.StatusNotFound:
		utils.APIResponse(context, "There's no book with specified id", http.StatusNotFound, http.MethodGet, err)
		return
	}

}

func (h *bookHandler) GetByIdHandler(context *gin.Context) {
	idStr := context.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		logrus.Warnf("Malformed parameter. %s", err.Error())
		utils.ValidatorErrorResponse(context, http.StatusBadRequest, http.MethodGet, err.Error())
		return
	}

	result, statusCode, err := h.service.FindById(uint(id))

	switch statusCode {
	case http.StatusOK:
		utils.APIResponse(context, "Book has been retrieved successfully.", http.StatusOK, http.MethodGet, result)
		return

	case http.StatusNotFound:
		utils.APIResponse(context, fmt.Sprintf("Book with id %d not found", id), http.StatusNotFound, http.MethodGet, err)
		return
	}

}

func (h *bookHandler) Test(ginContext *gin.Context) {
	logrus.Warn("JWT START")
	conn, _ := grpc.Dial("jwt-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	defer conn.Close()

	c := jwt.NewJwtServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, _ := c.GenerateToken(ctx, &jwt.JwtRequest{
		Username: ginContext.Param("username"),
	})

	utils.APIResponse(ginContext, "JWT.", http.StatusOK, http.MethodGet, resp)
}
