"use strict";

const {PostService} = require('../services/PostService');
const postService = new PostService();

let postController = {
    list: async function (req, res) {
        res.json(await postService.getList());
    },

    create: async function (req, res) {
        res.json(await postService.create(req.body));
    },
};
module.exports = postController;