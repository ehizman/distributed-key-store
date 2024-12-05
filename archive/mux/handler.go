package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloMuxHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello from a web service built with Gorilla Mux")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", helloMuxHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}