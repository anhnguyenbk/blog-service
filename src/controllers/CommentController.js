"use strict";

const {CommentService} = require('../services/CommentService');
const commentService = new CommentService();

let commentController = {
    get: async function (req, res) {
        res.json(await commentService.get(req.params.postId));
    },

    add: async function (req, res) {
        res.json(await commentService.create(req.params.postId, req.body));
    },

    delete: async function (req, res) {
        await commentService.delete(req.params.postId, req.params.commentId);
        res.status(204).send();
    }
};
module.exports = commentController;