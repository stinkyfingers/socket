import CardActions from '../actions/card';
import Reflux from 'reflux';
import config from '../config';

var CardStore = Reflux.createStore({
	listenables: [CardActions],

	create: function(card, type) {
		let code = 0;
		const route = type === 'Dealer' ? '/dealer' : '/card';
		const url = config.api + route;
		fetch(url, {
			method: 'POST',
			body: JSON.stringify(card)
		}).then((resp) => {
			code = resp.status;
			return resp.json();
		}).then((resp) => {
			if (code !== 200) {
				this.trigger({ error: resp });
				return;
			}
			this.trigger({ card: resp });
		});
	},
	unreviewed: function() {
		let code = 0;
		const url = config.api + '/unreviewed';
		fetch(url, {
			method: 'GET'
		}).then((resp) => {
			code = resp.status;
			return resp.json();
		}).then((resp) => {
			if (code !== 200) {
				this.trigger({ error: resp });
				return;
			}
			this.trigger({ cards: resp.cards, dealerCards: resp.dealerCards });
		});
	},
	approve: function(card, type, user) {
		let code = 0;
		let path = type === 'card' ? '/card' : '/dealerCard';
		path = path + '/' + user._id;
		const url = config.api + path;
		fetch(url, {
			method: 'PUT',
			body: JSON.stringify(card)
		}).then((resp) => {
			code = resp.status;
			return resp.json();
		}).then((resp) => {
			if (code !== 200) {
				this.trigger({ error: resp });
				return;
			}
		});
	}

});

module.exports = CardStore;