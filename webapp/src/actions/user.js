import Reflux from 'reflux';

var UserActions = Reflux.createActions([
	'authenticate',
	'getUser',
	'unsetUser',
	'saveUser',
	'passwordReset'
]);

module.exports = UserActions;