const api = require('@opentelemetry/api');
const { NodeTracerProvider } = require('@opentelemetry/node');
const { SimpleSpanProcessor } = require('@opentelemetry/tracing');
const { JaegerExporter } = require('@opentelemetry/exporter-jaeger');
const { B3MultiPropagator } = require('@opentelemetry/propagator-b3');

const serviceName = process.env.SERVICE_NAME || 'bar-service';
const jaegerAgentHost =
  process.env.JAEGER_AGENT_HOST || 'simplest-agent.observability';

// this will throw warning could not load plugin @opentelemetry/plugin-express of module express. Error: Cannot find module '@opentelemetry/plugin-express'
// const provider = new NodeTracerProvider();

const provider = new NodeTracerProvider({
  plugins: {
    express: {
      enabled: false
    }
  }
});

api.propagation.setGlobalPropagator(new B3MultiPropagator());

provider.addSpanProcessor(
  new SimpleSpanProcessor(
    new JaegerExporter({
      serviceName: serviceName,
      host: jaegerAgentHost
    })
  )
);

provider.register();
