"use strict";

const postController = require('../controllers/PostController');
const userController = require('../controllers/UserController');

module.exports = function (app) {
    // Users
    app.post('/auth/token', userController.auth)

    // Posts
    app.get('/posts', postController.listAll);
    app.get('/posts/published', postController.listPublished);
    app.get('/posts/:id', postController.get);
    app.get('/posts/slug/:slug', postController.getBySlug);
    app.post('/posts', postController.create);
    app.put('/posts/:id', postController.update);
    app.delete('/posts/:id', postController.delete);
};