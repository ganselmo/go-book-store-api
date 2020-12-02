package chat

import (
	"net/http"

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

func NewHTTPtransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func makeEndpoints(s Service) []*endpoint {
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
