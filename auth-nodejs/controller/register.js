const { User } = require('../models')
const generator = require('generate-password')

class Register {
    static async registerUser(req, res) {
        try {
            console.log(req.body)
            const { name, phone, role } = req.body

            const password = generator.generate({
                length: 4,
                numbers: true
            })

            const findResult = await User.findOne( {
                where: {
                    name
                }
            } )
            if (findResult) {
                res.status(400).json({message:"name already exist!"})
            }

            const insertedData = await User.create({ 
                name, 
                password,
                role, 
                phone
             })

            res.status(200).json({message:insertedData})

        } catch (error) {
            res.status(500).json({
                message: "internal server error"
            })
        }
    }
}

module.exports = Register