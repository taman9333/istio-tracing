// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"go.opentelemetry.io/otel"

// 	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
// 	"go.opentelemetry.io/otel/exporters/trace/jaeger"

// 	sdktrace "go.opentelemetry.io/otel/sdk/trace"
// )

// var tr = otel.Tracer("xyz-service")

// func main() {
// 	initTracer()
// 	r := mux.NewRouter()
// 	r.Use(otelmux.Middleware("xyz-service"))
// 	r.HandleFunc("/", index)
// 	http.Handle("/", r)
// 	fmt.Println("Starting up on 3000")
// 	log.Fatal(http.ListenAndServe(":3000", nil))
// }

// func index(w http.ResponseWriter, req *http.Request) {
// 	resp, err := http.Get("http://baz-svc:3000/")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	sb := string(body)
// 	fmt.Println("XYZ service:", sb)
// 	if reqHeadersBytes, err := json.Marshal(req.Header); err != nil {
// 		log.Println("Could not Marshal Req Headers")
// 	} else {
// 		log.Println(string(reqHeadersBytes))
// 	}
// }

// func initTracer() {
// 	exporter, err := jaeger.NewRawExporter(jaeger.WithAgentEndpoint("simplest-agent.observability:6831"), jaeger.WithSDK(&sdktrace.Config{
// 		DefaultSampler: sdktrace.AlwaysSample(),
// 	}))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// cfg := sdktrace.Config{
// 	// 	DefaultSampler: sdktrace.AlwaysSample(),
// 	// }
// 	tp := sdktrace.NewTracerProvider(
// 		// sdktrace.WithConfig(cfg),
// 		sdktrace.WithSyncer(exporter),
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	otel.SetTracerProvider(tp)
// 	// otel.SetTextMapPropagator(propagation.TraceContext{})
// }

// // func initTracer() func() {
// // 	tp, fn, err := jaeger.NewExportPipeline(jaeger.WithAgentEndpoint("simplest-agent.observability:6831"), jaeger.WithSDK(&sdktrace.Config{
// // 		DefaultSampler: sdktrace.AlwaysSample(),
// // 	}))
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	// cfg := sdktrace.Config{
// // 	// 	DefaultSampler: sdktrace.AlwaysSample(),
// // 	// }
// // 	// tp := sdktrace.NewTracerProvider(
// // 	// 	// sdktrace.WithConfig(cfg),
// // 	// 	sdktrace.WithSyncer(exporter),
// // 	// )
// // 	// if err != nil {
// // 	// 	log.Fatal(err)
// // 	// }
// // 	otel.SetTracerProvider(tp)
// // 	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
// // 	return fn
// // }

// ---

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"go.opentelemetry.io/otel"

// 	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
// 	"go.opentelemetry.io/otel/exporters/trace/jaeger"
// 	"go.opentelemetry.io/otel/propagation"
// 	"go.opentelemetry.io/otel/trace"

// 	sdktrace "go.opentelemetry.io/otel/sdk/trace"
// )

// // var tr = otel.Tracer("xyz-service")

// func main() {
// 	tr, flush := initTracer()
// 	defer flush()
// 	r := mux.NewRouter()
// 	r.Use(otelmux.Middleware("XYZ", otelmux.WithTracerProvider(tr)))
// 	r.HandleFunc("/", index)
// 	http.Handle("/", r)
// 	fmt.Println("Starting up on 3000")
// 	log.Fatal(http.ListenAndServe(":3000", nil))
// }

// func index(w http.ResponseWriter, req *http.Request) {
// 	resp, err := http.Get("http://baz-svc:3000/")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	sb := string(body)
// 	fmt.Println("XYZ service:", sb)
// 	if reqHeadersBytes, err := json.Marshal(req.Header); err != nil {
// 		log.Println("Could not Marshal Req Headers")
// 	} else {
// 		log.Println(string(reqHeadersBytes))
// 	}
// }

// // func initTracer() {
// // 	exporter, err := jaeger.NewRawExporter(jaeger.WithAgentEndpoint("simplest-agent.observability:6831"), jaeger.WithSDK(&sdktrace.Config{
// // 		DefaultSampler: sdktrace.AlwaysSample(),
// // 	}))
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	// cfg := sdktrace.Config{
// // 	// 	DefaultSampler: sdktrace.AlwaysSample(),
// // 	// }
// // 	tp := sdktrace.NewTracerProvider(
// // 		// sdktrace.WithConfig(cfg),
// // 		sdktrace.WithSyncer(exporter),
// // 	)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	otel.SetTracerProvider(tp)
// // 	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
// // }

// // func initTracer() func() {
// // 	tp, fn, err := jaeger.NewExportPipeline(jaeger.WithAgentEndpoint("simplest-agent.observability:6831"), jaeger.WithSDK(&sdktrace.Config{
// // 		DefaultSampler: sdktrace.AlwaysSample(),
// // 	}))
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	// cfg := sdktrace.Config{
// // 	// 	DefaultSampler: sdktrace.AlwaysSample(),
// // 	// }
// // 	// tp := sdktrace.NewTracerProvider(
// // 	// 	// sdktrace.WithConfig(cfg),
// // 	// 	sdktrace.WithSyncer(exporter),
// // 	// )
// // 	// if err != nil {
// // 	// 	log.Fatal(err)
// // 	// }
// // 	otel.SetTracerProvider(tp)
// // 	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
// // 	return fn
// // }

// func initTracer() (trace.TracerProvider, func()) {
// 	// Create and install Jaeger export pipeline.
// 	tr, flush, err := jaeger.NewExportPipeline(
// 		jaeger.WithAgentEndpoint("simplest-agent.observability:6831"),
// 		jaeger.WithProcess(jaeger.Process{
// 			ServiceName: "XYZ",
// 		}),
// 		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
// 	)
// 	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return tr, flush
// }
