package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/propagation"
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
	fmt.Println("Starting up on 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	// fmt.Println("!!! ", req.Context().Value("X-B3-Traceid"))
	reqExt, _ := http.NewRequestWithContext(req.Context(), "GET", "http://baz-svc:3000/", nil)

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
	// cfg := sdktrace.Config{
	// 	DefaultSampler: sdktrace.AlwaysSample(),
	// }
	// tp := sdktrace.NewTracerProvider(
	// 	sdktrace.WithConfig(cfg),
	// 	sdktrace.WithSyncer(exporter),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	otel.SetTracerProvider(tr)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tr, flush
}

// package main

// import (
// 	"log"
// 	"net/http"
// 	"xyz/tracer"

// 	"github.com/gin-gonic/gin"
// 	gintrace "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin"
// 	option "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin"
// 	"go.opentelemetry.io/otel/api/global"
// )

// var tr = global.Tracer("jaeger-tracing-go-service")

// func main() {
// 	fn := tracer.InitJaeger()
// 	defer fn()

// 	// Init Router
// 	router := gin.Default()
// 	router.Use(gintrace.Middleware("jaeger-tracing-go-service", option.WithTracer(tr)))

// 	// Route Handlers / Endpoints
// 	router.GET("/hello", func(c *gin.Context) {
// 		c.String(http.StatusOK, "hello world")
// 	})

// 	log.Fatal(router.Run(":3000"))
// }
