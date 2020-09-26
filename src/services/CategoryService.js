const {Connection} = require('../db/index');
const { v4: uuidv4 } = require('uuid');

class   CategoryService {
    constructor() {
        this.conn = new Connection();
    }

    async getAll() {
        const categories = await this.conn.getCategoryCollection();
        const cursor = categories.find();
        return await cursor.toArray();
    }

    async getBySlug(slug) {
        const categories = await this.conn.getCategoryCollection();
        return await categories.findOne({slug: slug});
    }

    async findOneWithPosts(query) {
        console.log(`find one category with posts by query: ${JSON.stringify(query)}`);

        const categories = await this.conn.getCategoryCollection();
        const cursor = categories.aggregate([
            {"$match": query },
            { 
                "$lookup": {
                    from: 'posts',
                    localField: '_id',
                    foreignField: 'categories',
                    as: 'posts'
                }
            }
        ])
        return (await cursor.toArray())[0]
    }
};

module.exports = {
    CategoryService
}