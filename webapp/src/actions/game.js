import Reflux from 'reflux';

var GameActions = Reflux.createActions([
	'createGame',
	'addPlayer',
	'initGame',
	'getGame',
	'getGameFromStorage',
	'joinGame',
	'play',
	'connect'
]);

module.exports = GameActions;