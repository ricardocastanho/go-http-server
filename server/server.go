package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ServerService interface {
	// Start the server
	Listen(port string, router *mux.Router)
}

type app struct {}

func (a *app) Listen(port string, router *mux.Router) {
	fmt.Printf("🚀  Server ready at http://localhost%s\n", port)
  log.Fatal(http.ListenAndServe(port, router))
}

func NewApp() ServerService {
	return &app{}
}
