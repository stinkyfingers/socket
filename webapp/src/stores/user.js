import UserActions from '../actions/user';
import Reflux from 'reflux';
import config from '../config';

var UserStore = Reflux.createStore({
	listenables: [UserActions],

	authenticate: function(user) {
		// use john test
		let code = 0;
		const url = config.api + '/auth'
		fetch(url, {
			method: 'POST',
			body: JSON.stringify(user)
		}).then((resp) => {
			code = resp.status;
			return resp.json();
		}).then((resp) => {
			if (code !== 200) {
				this.trigger({ error: resp });
				return
			}
			this.storeUser(resp)
			this.trigger({ user: resp });
		});
	},

	storeUser: function(user) {
		sessionStorage.setItem('user', JSON.stringify(user));
	},

	getUser: function() {
		const user = sessionStorage.getItem('user');
		this.trigger({ user:JSON.parse(user) });
	},

	unsetUser: function() {
		sessionStorage.removeItem('user');
		this.trigger({ user: undefined });
	}
});

module.exports = UserStore;