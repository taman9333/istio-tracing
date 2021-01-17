'use strict';

const { LogLevel } = require('@opentelemetry/core');
const { NodeTracerProvider } = require('@opentelemetry/node');
// const { B3Propagator } = require('@opentelemetry/core');

const provider = new NodeTracerProvider({
  logLevel: LogLevel.ERROR
});

provider.register();
// provider.register({ propagator: new B3Propagator() });
