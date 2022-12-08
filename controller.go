package main

import (
	"github.com/valyala/fasthttp"
)

func ParseRequest(ctx *fasthttp.RequestCtx) (string, string) {
	return ctx.UserValue("username").(string), string(ctx.QueryArgs().Peek("password"))
}

func Verify(ctx *fasthttp.RequestCtx) {
	username, password := ParseRequest(ctx)
	hash := GetHash(&username)

	if hash == nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	} else if VerifyPassword(&hash, &password) {
		ctx.SetStatusCode(fasthttp.StatusOK)
	} else {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
	}
}

func SignUp(ctx *fasthttp.RequestCtx) {
	username, password := ParseRequest(ctx)

	passwordHash := HashPassword(&password)

	Insert(&username, &passwordHash)
}

func Reset(ctx *fasthttp.RequestCtx) {
	username, password := ParseRequest(ctx)
	passwordHash := HashPassword(&password)

	Update(&username, &passwordHash)
}
