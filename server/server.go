package server

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (server *Server) Start() {
	instance := echo.New()

	instance.GET("/", homeHandler)

	instance.GET("/examples", examplesHandler)
	instance.GET("/examples/:name", exampleHandler)

	instance.POST("/examples/counter/increment", incrementCounterHandler)
	instance.POST("/examples/counter/decrement", decrementCounterHandler)

	instance.Logger.Fatal(instance.Start(server.listenAddr))
}
