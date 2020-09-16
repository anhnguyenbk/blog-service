"use strict";

const postController = require('../controllers/PostController');
const userController = require('../controllers/UserController');

module.exports = function (app) {
    // Users
    app.post('/auth/token', userController.auth)

    // Posts
    app.get('/posts', postController.list);
    app.get('/posts/:id', postController.get);
    app.post('/posts', postController.create);
    app.put('/posts/:id', postController.update);
    app.delete('/posts/:id', postController.delete);
};