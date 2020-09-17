'use strict'

const express = require('express')
const bodyParser = require('body-parser')
const cors = require('cors')

const app = express()

// parse application/json
app.use(bodyParser.json())
app.use(cors())

const appRoutes = require('./src/routes/index')(app)

// Export your Express configuration so that it can be consumed by the Lambda handler
module.exports = app
