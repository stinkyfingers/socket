module.exports = {
	api: process.env.NODE_ENV === 'production' ? 'https://differencebetweenapi.herokuapp.com/' : 'http://localhost:7000',
	websocket: process.env.NODE_ENV === 'production' ? 'ws://differencebetweenapi.herokuapp.com/' : 'ws://localhost:7000'
}