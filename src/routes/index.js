"use strict";

const postController = require('../controllers/PostController');
const userController = require('../controllers/UserController');
const {permit} = require('../middlewares/authorization');

module.exports = function (app) {
    // Users
    app.post('/auth/token', userController.auth)

    // Posts
    app.get('/posts', permit('ROLE_ADMIN'), postController.listAll);
    app.get('/posts/published', postController.listPublished);
    app.get('/posts/:id', postController.get);
    app.get('/posts/slug/:slug', postController.getBySlug);
    app.post('/posts', permit('ROLE_ADMIN'), postController.create);
    app.put('/posts/:id', permit('ROLE_ADMIN'), postController.update);
    app.delete('/posts/:id', permit('ROLE_ADMIN'), postController.delete);
};