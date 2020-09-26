"use strict";

const {CategoryService} = require('../services/CategoryService');
const categoryService = new CategoryService();

let categoryController = {
    getAll: async function (req, res) {
        res.json(await categoryService.getAll());
    },

    getBySlug: async function (req, res) {
        res.json(await categoryService.getBySlug(req.params.slug));
    }
};
module.exports = categoryController;