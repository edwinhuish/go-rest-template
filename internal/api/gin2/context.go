package gin2

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

type HandlerFunc func(ctx *Context)

func Cover(h HandlerFunc) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := &Context{
			c,
		}

		h(ctx)
	}
}

func (ctx *Context) Resp() *JsonResult {
	return &JsonResult{
		ctx:  ctx,
		Code: http.StatusOK,
		Msg:  "ok",
		Data: nil,
	}
}
