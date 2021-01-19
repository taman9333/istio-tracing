const { LogLevel } = require('@opentelemetry/core');
const { NodeTracerProvider } = require('@opentelemetry/node');
const { SimpleSpanProcessor } = require('@opentelemetry/tracing');
const { JaegerExporter } = require('@opentelemetry/exporter-jaeger');
const { B3Propagator } = require('@opentelemetry/propagator-b3');

const serviceName = process.env.SERVICE_NAME || 'foo-service';
const jaegerAgentHost =
  process.env.JAEGER_AGENT_HOST || 'simplest-agent.observability';
// const jaegerAgentPort = process.env.JAEGER_AGENT_PORT || '6831';

const provider = new NodeTracerProvider({
  plugins: {
    http: {
      enabled: true,
      path: '@opentelemetry/plugin-http'
    }
  },
  express: {
    enabled: true,
    path: '@opentelemetry/plugin-express'
  },
  logLevel: LogLevel.ERROR
});

provider.register({ propagator: new B3Propagator() });

provider.addSpanProcessor(
  new SimpleSpanProcessor(
    new JaegerExporter({
      serviceName: serviceName,
      host: jaegerAgentHost
      // port: jaegerAgentPort
    })
  )
);

console.log('Tracing initialized');
