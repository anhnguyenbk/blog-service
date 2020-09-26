// middleware for authentication

var jwt = require('jsonwebtoken');
const config = require('config');

const signKey = config.get('signKey');

async function authorize(req, res, next) {
    if (req.originalUrl == '/auth/token') {
        next()
        return;
    }

    const authHeader = req.headers.authorization;
    if (!authHeader) {
        res.status(401).json({message: "Unauthorized"});
    }

    const token = authHeader.split(' ')[1];
    // verify token
    jwt.verify(token, signKey, function(err, decoded) {
        if (err != null) {
            res.status(401).json({message: "Unauthorized"});
        } else {
            console.log (decoded);
            req.user = decoded;
            next();
        }
    });
}

module.exports = authorize;