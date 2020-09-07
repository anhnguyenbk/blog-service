"use strict";

const {MongoClient} = require('mongodb');
const config = require('config');

const dbConfig = config.get('dbConfig');

var db;

module.exports = new Promise(function(resolve, reject) {
    if (db == undefined) {
        console.log("Connect to " + dbConfig.uri);
        MongoClient.connect(dbConfig.uri, { useNewUrlParser: true, poolSize: dbConfig.poolSize })
            .then(client => {
                db = client.db(dbConfig.dbName)
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