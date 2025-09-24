package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *httpServer) handleIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := "Hello, world!"
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(response))
}
