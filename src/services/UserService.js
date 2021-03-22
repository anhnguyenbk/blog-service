const {Connection} = require('../db/index');
const { v4: uuidv4 } = require('uuid');
var jwt = require('jsonwebtoken');
var bcrypt = require('bcryptjs');
const config = require('config');

const signKey = config.get('signKey');

class UserService {
    constructor() {
        this.conn = new Connection();
    }

    async auth(email, password) {
        var collection = await this.conn.getUserCollection();

        var query = { email: email };
        const user = await collection.findOne(query);
        const match = await bcrypt.compare(password, user.password);
     
        if (match) {
            const token = jwt.sign({ 
                username: user.username, 
                email: user.email,
                firstName: user.firstName,
                lastName: user.lastName,
                roles: user.roles
            }, signKey,
            { 
                expiresIn: '2 days'
            });
            return {token : token}
        }
        return null;
    }
}

module.exports = {UserService}