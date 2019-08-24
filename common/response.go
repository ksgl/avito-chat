package common

import (
	"github.com/valyala/fasthttp"
)

var (
	AlreadyExists = []byte(`{"reason":"user already exists"}`)
	DoesNotExist  = []byte(`{"reason":"user or/and chat doesn't exist"}`)
)

func WriteResponse(ctx *fasthttp.RequestCtx, status int, payload []byte) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(status)
	ctx.SetBody(payload)
}
