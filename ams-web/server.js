const express = require('express')
const app = express()

app.use(express.static('./build'))
app.use(express.static('./public'))

app.listen(process.env.PORT, () => {
    console.log(`server start at port ${process.env.PORT}`)
})