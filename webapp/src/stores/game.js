import GameActions from '../actions/game';
import Reflux from 'reflux';
import config from '../config';

var GameStore = Reflux.createStore({
	listenables: [GameActions],
	ws: null,

	createGame: function(user) {
		let code = 0;
		const url = config.api + '/game/new';
		fetch(url, {
			method: 'POST',
			body: JSON.stringify(user)
		}).then((resp) => {
			code = resp.status;
			return resp.json();
		}).then((resp) => {
			if (code !== 200) {
				this.trigger({ error: resp });
				return;
			}
			this.storeGame(resp);
			this.trigger({ game: resp });
		});
	},

	initGame: function(game) {
		let code = 0;
		const url = config.api + '/game/init/' + game._id;
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
			this.storeGame(resp);
			this.trigger({ game: resp });
		});
	},

	getGame: function(id) {
		let code = 0;
		const url = config.api + '/game/' + id;
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
			this.storeGame(resp);
			this.trigger({ game: resp });
		});
	},

	joinGame: function(game, user) {
		let code = 0;
		const url = config.api + '/game/add/' + game._id;
		fetch(url, {
			method: 'POST',
			body: JSON.stringify(user)
		}).then((resp) => {
			code = resp.status;
			return resp.json();
		}).then((resp) => {
			if (code !== 200) {
				this.trigger({ error: resp });
				return;
			}
			this.storeGame(resp);
			this.trigger({ game: resp });
		});
	},

	exitGame: function(game) {
		let code = 0;
		const url = config.api + '/game/exit/' + game._id;
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
			this.unsetGame(resp);
			this.trigger({ game: null });
		});
	},

	updateGame: function(game) {
		let code = 0;
		const url = config.api + '/game/update';
		fetch(url, {
			method: 'PUT',
			body: JSON.stringify(game)
		}).then((resp) => {
			code = resp.status;
			return resp.json();
		}).then((resp) => {
			if (code !== 200) {
				this.trigger({ error: resp });
				return;
			}
			this.unsetGame(resp);
			this.trigger({ game: null });
		});
	},

	play: function(id, user, play) {
		const ws = this.ws;
		ws.send(JSON.stringify(play));

		ws.onmessage = ((msg) => {
			const game = JSON.parse(msg.data);

			this.storeGame(game);
			this.trigger({ game: game }); // was gameUpdate (app.js)
		})
	},

	connect: function(id) {
		const ws = new WebSocket(config.websocket + '/play/' + id);
		ws.onopen = (() => {
			this.ws = ws;
		});
		ws.onmessage = ((msg) => {
			const game = JSON.parse(msg.data);

			this.storeGame(game);
			this.trigger({ game: game });
		});
	},

	storeGame: function(game) {
		sessionStorage.setItem('game', JSON.stringify(game));
	},

	getGameFromStorage: function() {
		const game = sessionStorage.getItem('game');
		this.trigger({ game: JSON.parse(game) });
	},

	unsetGame: function() {
		sessionStorage.removeItem('game');
		this.trigger({ game: undefined });
	},

	storeWS: function(ws) {
		sessionStorage.setItem('ws', JSON.stringify(ws));
	},

	getWS: function() {
		return JSON.parse(sessionStorage.getItem('ws'));
	}
});

module.exports = GameStore;