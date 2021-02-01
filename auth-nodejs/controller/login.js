const { User } = require('../models')
const jwt = require('jsonwebtoken')

class Login {
    static async LoginController(req, res) {
        try {
            const { phone, password } = req.body 
            
            const resultData =  await User.findOne({
                where: {
                    phone, 
                    password
                }
            })

            if (!resultData) {
                res.status(500).json({message: "phone and password combination is not matched"})
            }

            const payload = {
                name: resultData.name,
                role: resultData.role, 
                phone: resultData.phone,
                timestamp: Date.now()
            }

            const token = jwt.sign(payload, process.env.SECRET, {algorithm: "HS512"})
            
            res.status(200).json({token})

        } catch (error) {
            res.status(500).json({message: error})
        }
    } 
}

module.exports = Login