'use strict'

const express = require('express')
const bodyParser = require('body-parser')
const cors = require('cors')
const authenticate = require('./src/middlewares/authentication')

const app = express()

app.use(bodyParser.json())
app.use(cors())
app.use(authenticate)

const appRoutes = require('./src/routes/index')(app)

// Export your Express configuration so that it can be consumed by the Lambda handler
module.exports = app
