package server

import (
	"github.com/julienschmidt/httprouter"
)

func SetupRouter() *httprouter.Router {
	router := httprouter.New()
	server := New()

	router.GET("/hello/:name", server.handleHello)
	router.GET("/", server.handleIndex)
	return router
}
