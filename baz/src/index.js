require('./tracing');
const express = require('express');
const app = express();
const port = 3000;

app.get('/', (req, res) => {
  console.log('Someone call Baz service');
  console.log(`## ${JSON.stringify(req.headers)}`);
  res.send('Hello from Baz service!');
});

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
});
