package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *httpServer) handleIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tag := "v0.1.0"
	response := fmt.Sprintf("Hello, world! %v", tag)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(response))
}
