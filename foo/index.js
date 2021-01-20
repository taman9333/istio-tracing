const express = require('express');
const app = express();
const port = 3000;
const axios = require('axios');

app.get('/', async (req, res) => {
  // calling Bar service & log response
  // Forward Headers for tracing
  // const headers = traceHeaders(req);

  // console.log(`## ${JSON.stringify(headers)}`);
  console.log(`## ${JSON.stringify(req.headers)}`);
  const barRes = await axios.get('http://bar-svc:3000/');
  await axios.get('http://baz-svc:3000/');

  // console.log(`$$ ${JSON.stringify(barRes.data)}`);

  res.send('Hello from Foo service!');
});

// function traceHeaders(req) {
//   incoming_headers = [
//     'x-request-id',
//     'x-b3-traceid',
//     'x-b3-spanid',
//     'x-b3-parentspanid',
//     'x-b3-sampled',
//     'x-b3-flags',
//     'x-ot-span-context'
//   ];
//   const headers = {};
//   for (let h of incoming_headers) {
//     if (req.header(h)) headers[h] = req.header(h);
//   }
//   return headers;
// }

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
});
