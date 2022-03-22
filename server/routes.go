package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", sayHello).Methods(http.MethodGet)

	return router
}
