const express = require('express')
const app = express()
const port = 3000
const axios = require('axios')

app.get('/', async (req, res) => {
  // calling Bar service & log response
  const barRes = await axios.get("http://bar-svc:3000/")
  console.log(`$$ ${JSON.stringify(barRes.data)}`)

  res.send('Hello from Foo service!')
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})