package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *httpServer) handleHello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username := ps.ByName("name")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Hello %v, u'r welcome!", username)

	if _, err := w.Write([]byte(response)); err != nil {
		log.Printf("error writing hello to user %v", username)
	}
}
