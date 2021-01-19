const { LogLevel } = require('@opentelemetry/core');
const { NodeTracerProvider } = require('@opentelemetry/node');
const { B3Propagator } = require('@opentelemetry/propagator-b3');

const provider = new NodeTracerProvider({
  logLevel: LogLevel.ERROR
});

provider.register({ propagator: new B3Propagator() });
