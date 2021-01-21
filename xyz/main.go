package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"xyz/ocmux"

	"github.com/gorilla/mux"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/plugin/ochttp/propagation/tracecontext"
)

func main() {
	ocmux.InitOpenCensusWithJaeger()
	r := mux.NewRouter()
	r.Use(ocmux.Middleware())
	r.HandleFunc("/hello", index)
	http.Handle("/", r)
	var handler http.Handler = r
	handler = &ochttp.Handler{
		Handler:     handler,
		Propagation: &tracecontext.HTTPFormat{}}

	fmt.Println("Starting up on 3000")
	log.Fatal(http.ListenAndServe(":3000", handler))
	// fmt.Println("sd")
}

func index(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("http://baz-svc:3000/")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	fmt.Println("XYZ service:", sb)
	if reqHeadersBytes, err := json.Marshal(req.Header); err != nil {
		log.Println("Could not Marshal Req Headers")
	} else {
		log.Println(string(reqHeadersBytes))
	}
	fmt.Fprintln(w, "Hello world!")
}
