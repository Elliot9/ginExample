package context

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Context interface {
	Method() string
	Host() string
	URI() string
	JSON(code int, messages any)
	HTML(name string, obj any)
	Redirect(code int, location string)
	Abort(code int, err error)
	// 反序列化	queryString
	ShouldBindQuery(obj any) error
	// 反序列化 PostForm
	// tags: `form:"id"`
	ShouldBindForm(obj any) error
	// 反序列化 Path Params (ex. /user/:id)
	// tags: `uri: "id"`
	ShouldBindURI(obj any) error
	// 反序列化 PostJson
	// tags: `json:"id"`
	ShouldBindJson(obj any) error
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

func (c *context) Redirect(code int, location string) {
	c.ctx.Redirect(code, location)
}

func (c *context) Abort(code int, err error) {
	c.ctx.AbortWithError(code, err)
}

func (c *context) ShouldBindQuery(obj any) error {
	return c.ctx.ShouldBindQuery(obj)
}

func (c *context) ShouldBindForm(obj any) error {
	return c.ctx.ShouldBindWith(obj, binding.FormPost)
}

func (c *context) ShouldBindURI(obj any) error {
	return c.ctx.ShouldBindUri(obj)
}

func (c *context) ShouldBindJson(obj any) error {
	return c.ctx.ShouldBindBodyWithJSON(obj)
}
