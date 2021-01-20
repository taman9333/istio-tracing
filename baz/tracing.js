const { LogLevel } = require('@opentelemetry/core');
const api = require('@opentelemetry/api');
const { NodeTracerProvider } = require('@opentelemetry/node');
const { SimpleSpanProcessor } = require('@opentelemetry/tracing');
const { JaegerExporter } = require('@opentelemetry/exporter-jaeger');
const { B3MultiPropagator } = require('@opentelemetry/propagator-b3');

const serviceName = process.env.SERVICE_NAME || 'baz-service';
const jaegerAgentHost =
  process.env.JAEGER_AGENT_HOST || 'simplest-agent.observability';
// const jaegerAgentPort = process.env.JAEGER_AGENT_PORT || '6831';

const provider = new NodeTracerProvider();

api.propagation.setGlobalPropagator(new B3MultiPropagator());

provider.addSpanProcessor(
  new SimpleSpanProcessor(
    new JaegerExporter({
      serviceName: serviceName,
      host: jaegerAgentHost
      // port: jaegerAgentPort
    })
  )
);

provider.register();
