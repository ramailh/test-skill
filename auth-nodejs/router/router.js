const {Register, Login, VerifyController} = require('../controller')
const express = require('express')
const jwt = require('../middlewares/jwt')
const router = express.Router()

router.post("/register", Register.registerUser)
router.post("/login", Login.LoginController)
router.use(jwt)
router.get("/verify", VerifyController.verify)

module.exports = router
