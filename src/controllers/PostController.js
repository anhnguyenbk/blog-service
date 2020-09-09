"use strict";

const {PostService} = require('../services/PostService');
const postService = new PostService();

let postController = {
    list: async function (req, res) {
        res.json(await postService.list(req));
    },

    get: async function (req, res) {
        res.json(await postService.get(req.params.id));
    },

    create: async function (req, res) {
        res.json(await postService.create(req.body));
    },

    update: async function (req, res) {
        res.json(await postService.update(req.params.id, req.body));
    },

    delete: async function (req, res) {
        await postService.delete(req.params.id);
        res.status(204).send();
    },
};
module.exports = postController;