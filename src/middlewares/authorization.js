// middleware for doing role-based permissions
function permit(...permittedRoles) {
    // return a middleware
    return (request, response, next) => {
      const { user } = request
  
      if (user && hasRole(user, permittedRoles)) {
        next(); // role is allowed, so continue on the next middleware
      } else {
        response.status(403).json({message: "Forbidden"}); // user is forbidden
      }
    }
}

function hasRole(user, permittedRoles) {
    for (const role of user.roles) {
      if (permittedRoles.includes(role)) {
        return true;
      }
    }
    return false;
}

module.exports = {permit}