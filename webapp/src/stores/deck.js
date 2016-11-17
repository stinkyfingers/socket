import DeckActions from '../actions/deck';
import Reflux from 'reflux';
import config from '../config';

var DeckStore = Reflux.createStore({
	listenables: [DeckActions],

	allDecks: function() {
		let code = 0;
		const url = config.api + '/decks';
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
			this.trigger({ decks: resp });
		});
	},

	getDeck: function(id) {
		let code = 0;
		const url = config.api + '/deck/' + id;
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
			this.trigger({ deck: resp });
		});
	},

	saveDeck: function(deck) {
		let code = 0;
		const url = config.api + '/deck';
		fetch(url, {
			method: 'PUT',
			body: JSON.stringify(deck)
		}).then((resp) => {
			code = resp.status;
			return resp.json();
		}).then((resp) => {
			if (code !== 200) {
				this.trigger({ error: resp });
				return;
			}
			this.trigger({ deck: resp });
		});
	}
});

module.exports = DeckStore;