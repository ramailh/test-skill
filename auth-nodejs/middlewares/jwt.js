const jwt = require('jsonwebtoken')

module.exports = async (req, res, next) => {
    try {
        const { authorization } = req.headers

        const claims = jwt.verify(authorization, process.env.SECRET, {algorithm: "HS512"})
        if (!claims) {
            res.status(401).json({message:"unauthorize token"})
        } 
        
        req.claims = claims
        next()
    } catch (error) {
        res.status(500).json({message:"internal server error"})
    }
}