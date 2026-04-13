package tests

import (
	"net"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func CreateServer() *httptest.Server {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/success", successHandler())
	r.PUT("/create", createHandler())

	listener, err := net.Listen("tcp", "127.0.0.1:3001")
	if err != nil {
		panic(err)
	}

	server := httptest.NewUnstartedServer(r)
	server.Listener = listener
	server.Start()

	return server
}

func successHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "everything is ok")
	}
}

func createHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusCreated, "created")
	}
}
