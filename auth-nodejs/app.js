require('dotenv').config()

const express = require('express')
const app = express()
const router = require('./router/router')


app.use(express.json())
app.use(router)

app.listen(process.env.PORT, () => {
    console.log("listening port on " + process.env.PORT)
})
