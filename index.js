const express = require('express')
var bodyParser = require('body-parser')

const app = express()
const port = 3000
// parse application/json
app.use(bodyParser.json())

const appRoutes = require('./src/routes/index')(app)
//const db = require('./src/dbs/index')

var server = app.listen(port, () => {
    console.log(`Server listening at http://localhost:${port}`);
});
