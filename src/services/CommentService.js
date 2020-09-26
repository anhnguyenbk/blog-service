const {Connection} = require('../db/index');
const { v4: uuidv4 } = require('uuid');

class CommentService {
    constructor() {
        this.conn = new Connection();
    }

    async get(postId) {
        const posts = await this.conn.getPostCollection();
        console.log(`Get comment by postId: ${postId}`);
        const post = await posts.findOne({ _id: postId });
        return post.comments;
    }

    async create(postId, comment) {
        const posts = await this.conn.getPostCollection();

        // find by document id and update and push item in array
        comment.id = uuidv4();
        comment.createdAt = Date.now();

        var result = await posts.findOneAndUpdate({ _id: postId },
            {$push: {comments: comment}}
        );
        return comment;
    }

    async delete(postId, commentId) {
        const posts = await this.conn.getPostCollection();

        await posts.findOneAndUpdate({ _id: postId },
            { $pull: { comments: { id: commentId } }},
        );
    }
};

module.exports = {
    CommentService
}