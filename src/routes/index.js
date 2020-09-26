"use strict";

const postController = require('../controllers/PostController');
const userController = require('../controllers/UserController');
const commentController = require('../controllers/CommentController');
const categoryController = require('../controllers/CategoryController');

const {permit} = require('../middlewares/authorization');

module.exports = function (app) {
    // Users
    app.post('/auth/token', userController.auth)

    // Posts
    app.get('/posts', permit('ROLE_ADMIN'), postController.listAll);
    app.get('/posts/published', postController.listPublished);
    app.get('/posts/:id', postController.get);
    app.get('/posts/slug/:slug', postController.getBySlug);
    app.get('/posts/category/:id', postController.getByCategory);
    app.post('/posts', permit('ROLE_ADMIN'), postController.create);
    app.put('/posts/:id', permit('ROLE_ADMIN'), postController.update);
    app.delete('/posts/:id', permit('ROLE_ADMIN'), postController.delete);

    // Comments
    app.get('/posts/:postId/comments', commentController.get);
    app.post('/posts/:postId/comments', commentController.add);
    app.delete('/posts/:postId/comments/:commentId', permit('ROLE_ADMIN'), commentController.delete);

    // Categories
    app.get('/categories', categoryController.getAll);
    app.get('/categories/slug/:slug', categoryController.getBySlug);
};