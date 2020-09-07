'use strict'

const express = require('express')
var bodyParser = require('body-parser')
var cors = require('body-parser')

const app = express()
const port = 3000
// parse application/json
app.use(bodyParser.json())
app.use(cors())

const appRoutes = require('./src/routes/index')(app)
// var server = app.listen(port, () => {
//     console.log(`Server listening at http://localhost:${port}`);
// });

// Export your Express configuration so that it can be consumed by the Lambda handler
module.exports = app
