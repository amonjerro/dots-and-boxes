require('dotenv').config()

const express = require('express')
const app = express()

app.use(express.static('static'))

app.listen(process.env.PORT, () =>{
    console.log('Happy hacking')
})