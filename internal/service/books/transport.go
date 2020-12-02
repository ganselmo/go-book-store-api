package book

import (
	"net/http"
	"strconv"

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

func (s httpService) Register(r *gin.Engine) {

	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}

func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/books",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/books/:id",
		function: getById(s),
	})

	list = append(list, &endpoint{
		method:   "POST",
		path:     "/books",
		function: saveBook(s),
	})

	list = append(list, &endpoint{
		method:   "PATCH",
		path:     "/books/:id",
		function: editBook(s),
	})

	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/books/:id",
		function: deleteBook(s),
	})

	return list
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"books:": s.FindAll(),
		})
	}
}

func getById(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		b := s.FindById(id)
		m := http.StatusOK
		if b == nil {
			m = http.StatusNotFound
		}
		c.JSON(m, gin.H{
			"book": b,
		})
	}
}

func saveBook(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var b Book
		err := c.BindJSON(&b)
		if err != nil {
			panic(err)
		}
		s.AddBook(b)

	}
}

func editBook(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var b Book
		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		b.Id = id
		err = c.BindJSON(&b)
		if err != nil {
			panic(err)
		}
		s.PatchBook(b)

	}
}

func deleteBook(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		s.DeleteBook(id)

	}
}
