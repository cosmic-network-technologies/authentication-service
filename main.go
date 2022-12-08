package main

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"time"
)

func main() {
	connect()

	r := router.New()

	r.GET("/verify/{username}", Verify)
	r.POST("/signup/{username}", SignUp)
	r.POST("/reset/{username}", Reset)

	server := &fasthttp.Server{
		Handler:           r.Handler,
		Name:              "authentication-service",
		ReadTimeout:       60 * time.Second,
		ReduceMemoryUsage: true,
	}

	err := server.ListenAndServe("127.0.0.1:8080")

	if err != nil {
		panic(err)
	}
}
