"use strict";

const {MongoClient} = require('mongodb');

const uri = "mongodb+srv://anhngnet:f0HQ7scGSqYUoPcK@cluster0.vbmqb.mongodb.net?retryWrites=true&w=majority";
const db_name = "anhngblog";

var db;

module.exports = new Promise(function(resolve, reject) {
    if (db == undefined) {
        console.log("Connect to " + uri);
        MongoClient.connect(uri, { useNewUrlParser: true, poolSize: 10 })
            .then(client => {
                db = client.db(db_name)
                resolve(db);
            })
            .catch(err => {
                console.error(err);
                reject (err)
            });
    } else {
        resolve(db);
    }
});