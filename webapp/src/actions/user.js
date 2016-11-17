import Reflux from 'reflux';

var UserActions = Reflux.createActions([
	'authenticate',
	'getUser',
	'unsetUser'
]);

module.exports = UserActions;