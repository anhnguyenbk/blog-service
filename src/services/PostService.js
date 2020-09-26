const {Connection} = require('../db/index');
const { v4: uuidv4 } = require('uuid');

class PostService {
    constructor() {
        this.conn = new Connection();
    }

    async listAll(req) {
        const query = {
            status: {
                $ne: 'deleted'
            }
        };
        return await this.findAll(query)
    }

    async listPublished(req) {
        const query = {
            status: {
                $eq: 'published'
            }
        };
        return await this.findAll(query);
    }

    async get(id) {
        return await this.findOne({ _id: id })
    }

    async getBySlug(slug) {
        return await this.findOne({ slug: slug });
    }

    async getByCategory(id) {
        return await this.findAll({ 
            categories: id, 
            status: {
                $eq: 'published'
            }
        });
    }

    async create(post) {
        const collection = await this.conn.getPostCollection();

        post._id = uuidv4();
        post.createdAt = Date.now();
        post.updatedAt = Date.now();

        const result = await collection.insertOne(post);
        console.log(`Post inserted with the _id: ${result.insertedId}`);
        return result.ops[0];
    }

    async update(id, post) {
        const collection = await this.conn.getPostCollection();

        var query = { _id: id };
        var values = { $set: { 
            title: post.title, 
            slug: post.slug, 
            desc: post.desc, 
            content: post.content, 
            updatedAt: Date.now() 
        } }
        const result = await collection.findOneAndUpdate(query, values, { returnOriginal: false });
        console.log(`Post updated with the _id: ${id}`);
        return result.value
    }

    async delete(id) {
        const collection = await this.conn.getPostCollection();
        var query = { _id: id };
        var values = { $set: {status: "deleted" } };
        await collection.updateOne(query, values);
    }

    async findOne(query) {
        const collection = await this.conn.getPostCollection();

        console.log(`findOne post by query: ${JSON.stringify(query)}`);
        const cursor = collection.aggregate([
            { "$match": query },
            {
                "$lookup": {
                    "from": "categories",
                    "localField": "categories",
                    "foreignField": "_id",
                    "as": "categories"
                }
            }
        ]);
        return (await cursor.toArray())[0];
    }

    async findAll(query) {
        console.log(`find all post by query: ${JSON.stringify(query)}`);

        const collection = await this.conn.getPostCollection();
        const cursor = collection.aggregate([
            { "$match": query },
            {
                "$lookup": {
                    "from": "categories",
                    "localField": "categories",
                    "foreignField": "_id",
                    "as": "categories"
                }
            }
        ]);
        return await cursor.toArray();
    }
};

module.exports = {
    PostService
}