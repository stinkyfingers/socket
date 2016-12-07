'use strict';
const express = require('express');
const path = require('path');
const app = express();

app.get('/player', (r, w) => {
	w.sendFile('build/index.html', {root: __dirname});
});
app.get('/create', (r, w) => {
	w.sendFile('build/index.html', {root: __dirname});
});
app.all('/play/*', (r, w) => {
	w.sendFile('build/index.html', {root: __dirname});
});
app.all('/card', (r, w) => {
	w.sendFile('build/index.html', {root: __dirname});
});
app.all('/audit', (r, w) => {
	w.sendFile('build/index.html', {root: __dirname});
});
app.all('/', (r, w) => {
	w.sendFile('build/index.html', {root: __dirname});
});

app.use(express.static(path.join(__dirname, 'build')));

const PORT = process.env.PORT || 8080;
app.listen(PORT, () => {
	console.log(`app running on ${PORT}`);
});