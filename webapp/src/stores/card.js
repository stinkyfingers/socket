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
				console.log(resp)
				return;
			}
			this.trigger({ card: resp });
		});
	}
});

module.exports = CardStore;