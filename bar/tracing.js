'use strict';

const { LogLevel } = require('@opentelemetry/core');
const { NodeTracerProvider } = require('@opentelemetry/node');
const { SimpleSpanProcessor } = require('@opentelemetry/tracing');
const { JaegerExporter } = require('@opentelemetry/exporter-jaeger');
// const { B3Propagator } = require('@opentelemetry/core');

const provider = new NodeTracerProvider({
  plugins: {
    express: {
      enabled: true,
      path: '@opentelemetry/plugin-express'
    },
    http: {
      enabled: true,
      path: '@opentelemetry/plugin-http'
    }
  },
  logLevel: LogLevel.ERROR
});

provider.addSpanProcessor(
  new SimpleSpanProcessor(
    new JaegerExporter({
      serviceName: 'bar-started',
      url: 'jaeger-collector.istio-system'
      // If you are running your tracing backend on another host,
      // you can point to it using the `url` parameter of the
      // exporter config.
    })
  )
);

provider.register();
// provider.register({ propagator: new B3Propagator() });
