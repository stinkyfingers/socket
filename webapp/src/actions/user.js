import Reflux from 'reflux';

var UserActions = Reflux.createActions([
	'authenticate',
	'getUser',
	'unsetUser',
	'saveUser'
]);

module.exports = UserActions;