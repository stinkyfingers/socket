import Reflux from 'reflux';

var CardActions = Reflux.createActions([
	'create',
	'unreviewed',
	'approve'
]);

module.exports = CardActions;