"use strict";

const {MongoClient} = require('mongodb');
const config = require('config');

const dbConfig = config.get('dbConfig');
var db;

class Connection {
    constructor(){
    }

    async getPostCollection() {
        var conn = await this.openConnection();
        return conn.collection("posts");
    }

    async getUserCollection() {
        var conn = await this.openConnection();
        return conn.collection("users");
    }

    async getCategoryCollection() {
        var conn = await this.openConnection();
        return conn.collection("categories");
    }

    openConnection() {
        return new Promise((resolve, reject) => {
            if (db == undefined) {
                console.log("Open connect to " + dbConfig.uri);
                MongoClient.connect(dbConfig.uri, { useNewUrlParser: true, poolSize: dbConfig.poolSize, useUnifiedTopology: true })
                    .then(client => {
                        console.log("Connection open successfully")
                        db = client.db(dbConfig.dbName)
                        resolve(db);
                    })
                    .catch(err => {
                        console.error("An error has happenned while connect to " + dbConfig.uri);
                        console.error(err);
                        reject (err)
                    });
            } else {
                resolve(db);
            }
        })
    } 
    
}

module.exports = {Connection}
