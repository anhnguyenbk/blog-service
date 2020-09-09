const {Connection} = require('../db/index');
const { v4: uuidv4 } = require('uuid');

class PostService {
    constructor() {
        this.conn = new Connection();
    }

    async list(req) {
        const collection = await this.conn.getPostCollection();

        const query = {
            status: {
                $eq: 'published'
            }
        };

        const options = {
          sort: { createdAt: 1 },
        };
        const cursor = collection.find(query, options);
        return await cursor.toArray();
    }

    async get(id) {
        const collection = await this.conn.getPostCollection();

        console.log(`Get post with the _id: ${id}`);
        var query = { _id: id };
        return await collection.findOne(query);
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
};

module.exports = {
    PostService
}