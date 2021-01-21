const tracing = require('@opencensus/nodejs');
const propagation = require('@opencensus/propagation-b3');
import { JaegerTraceExporter } from '@opencensus/exporter-jaeger';
const b3 = new propagation.B3Format();

const options = {
  serviceName: 'bar-service',
  host: 'simplest-agent.observability',
  port: 6831
};
const exporter = new JaegerTraceExporter(options);

tracing.start({
  exporter,
  propagation: b3,
  samplingRate: 1.0
});
