const { User } = require('../models')
const jwt = require('jsonwebtoken')

class VerifyController {
    static async verify(req, res) {
        try {
            const claims = req.claims
            res.status(200).json(claims)
        } catch (error) {
            res.status(500).json({message:"internal server error"})
        }
    }
}

module.exports = VerifyController