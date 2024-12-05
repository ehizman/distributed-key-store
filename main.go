package main

import (
	"log"
	"net/http"

	"github.com/ehizman/key-value-store/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/v1/key/{key}", handlers.PutHandlerFunc).Methods("PUT")
	r.HandleFunc("/v1/key/{key}", handlers.GetHandlerFunc).Methods("GET")
	r.HandleFunc("/v1/key/{key}", handlers.DeleteFuncHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
