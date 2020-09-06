"use strict";

const postController = require('../controllers/PostController');

module.exports = function (app) {
    app.get('/posts', postController.list);
    app.post('/posts', postController.create);
};