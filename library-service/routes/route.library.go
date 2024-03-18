package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	book_controllers "library-service/controllers/book-controllers"
	"library-service/handlers/book-handler"
	"library-service/middlewares"
)

func InitLibraryRoutes(db *gorm.DB, route *gin.RouterGroup) {
	bookRepository := book_controllers.NewBookRepository(db)
	bookService := book_controllers.NewBookService(bookRepository)
	bookHandler := book_handler.NewBookHandler(bookService)

	route.POST("/book", middlewares.Auth(), bookHandler.CreateHandler)
	route.GET("/book/:id", bookHandler.GetByIdHandler)
	route.GET("/book", bookHandler.GetAllHandler)
	route.DELETE("/book/:id", middlewares.Auth(), bookHandler.DeleteHandler)
}
