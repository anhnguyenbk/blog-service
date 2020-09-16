"use strict";

const {UserService} = require('../services/UserService');
const userService = new UserService();

let userController = {
    auth: async function (req, res) {
        const auth = await userService.auth(req.body.email, req.body.password);
        if (auth == null) {
            res.status(401).send({error: 'unauthorized'});
        } else {
            res.json(auth);
        }
    }
}

module.exports = userController;