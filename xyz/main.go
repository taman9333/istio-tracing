package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/label"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	tr, flush := initTracer()
	defer flush()
	r := mux.NewRouter()
	r.Use(otelmux.Middleware("XYZ", otelmux.WithTracerProvider(tr)))
	r.HandleFunc("/hello", index)
	http.Handle("/", r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("Starting up on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type myTransport struct {
	originalTransport http.RoundTripper
}

func (t *myTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Println("BEFORE REQ")
	if reqHeadersBytes, err := json.Marshal(req.Header); err != nil {
		log.Println("Could not Marshal Req Headers")
	} else {
		log.Println("***", string(reqHeadersBytes))
	}
	res, err := t.originalTransport.RoundTrip(req)
	fmt.Println("AFTER RES")
	if reqHeadersBytes, err := json.Marshal(req.Header); err != nil {
		log.Println("Could not Marshal Req Headers")
	} else {
		log.Println("***", string(reqHeadersBytes))
	}
	return res, err
}

func index(w http.ResponseWriter, req *http.Request) {
	client := http.Client{
		Transport: &myTransport{otelhttp.NewTransport(http.DefaultTransport)},
	}

	ctx := baggage.ContextWithValues(req.Context(), label.String("foo", "bar"))

	reqExt, _ := http.NewRequestWithContext(ctx, "GET", "http://baz-svc:3000/", nil)
	if reqHeadersBytes, err := json.Marshal(reqExt.Header); err != nil {
		log.Println("Could not Marshal Req Headers")
	} else {
		log.Println("***", string(reqHeadersBytes))
	}

	resp, err := client.Do(reqExt)
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

func initTracer() (trace.TracerProvider, func()) {
	tr, flush, err := jaeger.NewExportPipeline(
		jaeger.WithAgentEndpoint("simplest-agent.observability:6831"),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		jaeger.WithProcess(jaeger.Process{ServiceName: "XYZ"}),
	)
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(tr)
	b3 := b3.B3{
		InjectEncoding: b3.B3MultipleHeader | b3.B3SingleHeader,
	}
	otel.SetTextMapPropagator(b3)
	return tr, flush
}
