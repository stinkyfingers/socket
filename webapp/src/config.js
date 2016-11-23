module.exports = {
	api: process.env.NODE_ENV === 'production' ? 'http://104.197.74.147' : 'http://localhost:7000',
	websocket: process.env.NODE_ENV === 'production' ? 'ws://104.197.74.147' : 'ws://localhost:7000'
}