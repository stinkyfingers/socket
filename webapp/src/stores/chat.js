import ChatActions from '../actions/chat';
import Reflux from 'reflux';
import config from '../config';

var ChatStore = Reflux.createStore({
	listenables: [ChatActions],
	messages: [],

	connect: function(id) {
		const ws = new WebSocket(config.websocket + '/chat/' + id);
		ws.onopen = (() => {
			this.ws = ws;
		});
		ws.onmessage = ((msg) => {
			const message = JSON.parse(msg.data);

			this.messages.push({message});
			this.trigger({ messages: this.messages });
		});
	},

	send: function(id, message, user) {
		this.ws.send(JSON.stringify({text: message, playerName: user.name, id: id }));
	}

});

module.exports = ChatStore;