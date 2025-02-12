package context

import (
	"io"
	"net/http"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const AbortErrorName = "_abort_error_"
const SessionAuthKey = "user"

type Context interface {
	Header() http.Header
	Method() string
	Host() string
	URI() string
	JSON(code int, messages any)
	HTML(name string, obj any)
	Redirect(code int, location string)
	Abort(CustomError)
	GetAbort() CustomError
	ReturnBackWith(messages map[string]any)
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
	Body() []byte
	ResponseWriter() gin.ResponseWriter
	Session() sessions.Session
	GetFlash() (map[string]any, bool)
}

type HandlerFunc func(c Context)

// Wrapper
func NewContext(ctx *gin.Context) Context {
	session := sessions.Default(ctx)
	return &context{
		ctx:     ctx,
		session: session,
	}
}

type context struct {
	ctx     *gin.Context
	session sessions.Session
}

func (c *context) Header() http.Header {
	header := c.ctx.Request.Header

	clone := make(http.Header, len(header))
	for k, v := range header {
		value := make([]string, len(v))
		copy(value, v)

		clone[k] = value
	}
	return clone
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

func (c *context) Abort(err CustomError) {
	c.ctx.Set(AbortErrorName, err)
	c.ctx.AbortWithStatus(err.HTTPCode())
}

func (c *context) GetAbort() (err CustomError) {
	value, exists := c.ctx.Get(AbortErrorName)
	if !exists {
		return nil
	}
	return value.(CustomError)
}

func (c *context) ReturnBackWith(messages map[string]any) {
	referer := c.Header().Get("Referer")
	if referer == "" {
		referer = "/"
	}

	c.Session().AddFlash(messages)
	c.Session().Save()
	c.Redirect(http.StatusFound, referer)
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

func (c *context) Body() []byte {
	body, _ := io.ReadAll(c.ctx.Request.Body)
	return body
}

func (c *context) ResponseWriter() gin.ResponseWriter {
	return c.ctx.Writer
}

func (c *context) Session() sessions.Session {
	return c.session
}

func (c *context) GetFlash() (map[string]any, bool) {
	result := c.Session().Flashes()

	c.Session().Delete("_flash")
	c.Session().Save()

	if len(result) == 0 {
		return make(map[string]any), false
	}

	flashMap := make(map[string]any)
	for _, v := range result {
		if m, ok := v.(map[string]any); ok {
			flashMap = m
			break
		}
	}

	if len(flashMap) == 0 {
		return make(map[string]any), false
	}

	return flashMap, true
}
