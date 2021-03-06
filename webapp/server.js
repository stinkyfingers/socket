'use strict';
const express = require('express');
const path = require('path');
const app = express();

app.all('/[a-z]+', (r, w) => {
	w.sendFile(path.join(__dirname, 'build', 'index.html'));
});

app.use(express.static(path.join(__dirname, 'build')));

const PORT = process.env.PORT || 8080;
app.listen(PORT, () => {
	console.log(`app running on ${PORT}`);
});