const tracing = require('@opencensus/nodejs');
const propagation = require('@opencensus/propagation-b3');
const { JaegerTraceExporter } = require('@opencensus/exporter-jaeger');
const b3 = new propagation.B3Format();

const options = {
  serviceName: 'baz-service',
  host: 'simplest-agent.observability',
  port: 6832
};
const exporter = new JaegerTraceExporter(options);

tracing.start({
  exporter,
  propagation: b3,
  samplingRate: 1.0
});
