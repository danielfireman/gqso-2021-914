const express = require('express');
const { raizQuadrada } = require('./raiz');
const app = express();

app.get('/', (req, res) => {
    res.send('Hello World!');
})

app.get('/raiz/:op', raizQuadrada);

app.listen(3333, () => {
    console.log('Running in port ' + 3333);
})