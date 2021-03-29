package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Article struct {}

func NewArticle() Article  {
	return Article{}
}

func (a Article) Get(c *gin.Context)  {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
func (a Article) List(c *gin.Context)  {}
func (a Article) Create(c *gin.Context)  {}
func (a Article) Update(c *gin.Context)  {}
func (a Article) Delete(c *gin.Context)  {}