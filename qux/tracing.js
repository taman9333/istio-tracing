const { LogLevel } = require('@opentelemetry/core');
const { NodeTracerProvider } = require('@opentelemetry/node');
const api = require('@opentelemetry/api');
const { SimpleSpanProcessor } = require('@opentelemetry/tracing');
const { JaegerExporter } = require('@opentelemetry/exporter-jaeger');
// const { B3Propagator } = require('@opentelemetry/propagator-b3');
const { B3MultiPropagator } = require('@opentelemetry/propagator-b3');
// const { B3SinglePropagator } = require('@opentelemetry/propagator-b3');

const provider = new NodeTracerProvider({
  logLevel: LogLevel.ERROR
});

// api.propagation.setGlobalPropagator(new B3Propagator());

api.propagation.setGlobalPropagator(new B3MultiPropagator());

// api.propagation.setGlobalPropagator(new B3SinglePropagator());

// provider.register({ propagator: new B3Propagator() });

provider.register();

// provider.addSpanProcessor(
//   new SimpleSpanProcessor(
//     new JaegerExporter({
//       serviceName: 'qux',
//       host: 'tracing.istio-system'
//     })
//   )
// );
