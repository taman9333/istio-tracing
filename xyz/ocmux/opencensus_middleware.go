package ocmux

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.opencensus.io/trace"

	"contrib.go.opencensus.io/exporter/jaeger"
)

const (
	jaegerURL = "http://localhost:9411/api/v2/spans"
)

// NameFromGorillaMux inspects the HTTP Request to see if it can extract the
// followed HTTP route to generate the needed OpenCensus span name. If not
// possible it will default to the URL Path.
// This is typically used in a HTTP client when using Gorilla Mux as URL
// constructor.
func NameFromGorillaMux(router *mux.Router) func(*http.Request) string {
	return func(r *http.Request) string {
		var match mux.RouteMatch
		if router.Match(r, &match) {
			if name, err := match.Route.GetPathTemplate(); err == nil {
				return r.Method + " " + name
			}
		}
		// default to URL Path
		return r.URL.Path
	}
}

// Middleware holds a Gorilla Mux middleware to update the OpenCensus span name.
// This is typically used in a HTTP server using Gorilla Mux for routing.
func Middleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			span := trace.FromContext(req.Context())
			if span == nil {
				next.ServeHTTP(w, req)
				return
			}
			route := mux.CurrentRoute(req)
			if route == nil {
				next.ServeHTTP(w, req)
				return
			}
			if name, err := route.GetPathTemplate(); err == nil {
				span.SetName(req.Method + " " + name)
			}
			next.ServeHTTP(w, req)
		})
	}
}

func InitOpenCensusWithJaeger() {
	// // Always trace for this demo.
	// trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	// rep := zipkinhttp.NewReporter(jaegerURL)

	// localEndpoint, _ := jaeger.NewEndpoint(serviceName, hostPort)

	// exporter := oczipkin.NewExporter(rep, localEndpoint)
	// trace.RegisterExporter(exporter)

	// return rep
	exporter, err := jaeger.NewExporter(jaeger.Options{
		CollectorEndpoint: "http://simplest-agent.observability:6832",
		Process: jaeger.Process{
			ServiceName: "XYZ-service",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	trace.RegisterExporter(exporter)

	// For demoing purposes, always sample. In a production application, you should
	// configure this to a trace.ProbabilitySampler set at the desired
	// probability.
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}
