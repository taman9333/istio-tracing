const { LogLevel } = require('@opentelemetry/core');
const { NodeTracerProvider } = require('@opentelemetry/node');
// const api = require('@opentelemetry/api');
const { B3Propagator } = require('@opentelemetry/propagator-b3');

const provider = new NodeTracerProvider({
  logLevel: LogLevel.ERROR
});

// api.propagation.setGlobalPropagator(new B3Propagator());

provider.register({ propagator: new B3Propagator() });
