package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	http.Handle("/", r)
	fmt.Println("Starting up on 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello world")
}
