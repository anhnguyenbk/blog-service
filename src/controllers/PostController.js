"use strict";

const {PostService} = require('../services/PostService');
const postService = new PostService();

let postController = {
    listAll: async function (req, res) {
        res.json(await postService.listAll(req));
    },
    
    listPublished: async function (req, res) {
        res.json(await postService.listPublished(req));
    },

    get: async function (req, res) {
        res.json(await postService.get(req.params.id));
    },

    getBySlug: async function (req, res) {
        res.json(await postService.getBySlug(req.params.slug));
    },

    getByCategory: async function (req, res) {
        res.json(await postService.getByCategory(req.params.id));
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