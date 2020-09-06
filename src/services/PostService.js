var postCollection;
var dbs = require('../dbs/index').then(function (db) {
    postCollection = db.collection("posts")
});

class PostService {
    constructor() {
        //this.collection = db.collection("posts")
    }

    async getList() {
        const query = {
            status: {
                $eq: 'published'
            }
        };
        // // const options = {
        // //   // sort returned documents in ascending order by title (A->Z)
        // //   sort: { title: 1 },
        // // // Include only the `title` and `imdb` fields in each returned document
        // //   projection: { _id: 0, title: 1, imdb: 1 },
        // // };
        // console.log(postCollection)
        const cursor = postCollection.find(query);
        return await cursor.toArray();

        return []
    }

    async create(post) {
        post.createdAt = Date.now();
        post.updatedAt = Date.now();

        const result = await postCollection.insertOne(post);
        console.log(`Post inserted with the _id: ${result.insertedId}`);
        return result.ops[0];
    }
};

module.exports = {
    PostService
}