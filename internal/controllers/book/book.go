package book

import (
	"net/http"
	"strconv"

	"github.com/ganselmo/go-book-store-api/internal/service"
	"github.com/gin-gonic/gin"
)

// HTTPController ...
type HTTPController struct {
	router  *gin.Engine
	service *service.AuthorService
}

// NewHTTPController ...
func NewHTTPController(s *service.AuthorService) HTTPController {
	r := gin.Default()
	makeEndpoints(r, s) // registro los endpoints
	return HTTPController{r, s}
}

// Run ejecuta el controller
func (c *HTTPController) Run() {
	c.router.Run()
}

func makeEndpoints(r *gin.Engine, s *service.BookService) {
	r.GET("/books/:id", makeFindHandler(s))
}

func makeFindHandler(s *service.BookService) gin.HandlerFunc {
	// Fijate que aca devuelvo una funcion y no un valor
	return func(c *gin.Context) {
		i, _ := strconv.Atoi(c.Param("id"))

		v := (*s).FindBooksById((i))

		c.JSON(http.StatusOK, gin.H{
			"book": v,
		})
	}
}
