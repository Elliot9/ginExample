package context

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Context interface {
	Method() string
	Host() string
	URI() string
	JSON(code int, messages any)
	HTML(name string, obj any)
}

type HandlerFunc func(c Context)

// Wrapper
func NewContext(ctx *gin.Context) Context {
	return &context{ctx: ctx}
}

type context struct {
	ctx *gin.Context
}

func (c *context) Method() string {
	return c.ctx.Request.Method
}

func (c *context) Host() string {
	return c.ctx.Request.Host
}

func (c *context) URI() string {
	uri, _ := url.QueryUnescape(c.ctx.Request.RequestURI)
	return uri
}

func (c *context) JSON(code int, messages any) {
	c.ctx.JSON(code, messages)
}

func (c *context) HTML(name string, obj any) {
	c.ctx.HTML(http.StatusOK, name, obj)
}
