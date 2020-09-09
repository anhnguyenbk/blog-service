"use strict";

const postController = require('../controllers/PostController');

module.exports = function (app) {
    app.get('/posts', postController.list);
    app.get('/posts/:id', postController.get);
    app.post('/posts', postController.create);
    app.put('/posts/:id', postController.update);
    app.delete('/posts/:id', postController.delete);
};