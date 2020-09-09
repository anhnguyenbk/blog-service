'use strict'

const express = require('express')
var bodyParser = require('body-parser')
var cors = require('body-parser')

const app = express()

// parse application/json
app.use(bodyParser.json())
app.use(cors())

const appRoutes = require('./src/routes/index')(app)

// Export your Express configuration so that it can be consumed by the Lambda handler
module.exports = app
