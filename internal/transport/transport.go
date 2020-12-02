package chat

import (
	"net/http"

	"github.com/ganselmo/go-book-store-api/internal/service/author"
	"github.com/ganselmo/go-book-store-api/internal/service/book"
	"github.com/gin-gonic/gin"
)

//HTTPService...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

func NewHTTPtransport(sb book.Service, sa author.Service) HTTPService {
	endpoints := makeBookEndpoints(sa)

	return httpService{endpoints}
}

func makeBookEndpoints(s book.Service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/books",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method:   "POST",
		path:     "/book",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/authors",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method:   "POST",
		path:     "/author",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/author/books",
		function: getAuthorBooks(s),
	})

	return list
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages:": s.FindAll(),
		})
	}
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages:": s.FindAll(),
		})
	}
}

func (s httpService) Register(r *gin.Engine) {

	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
