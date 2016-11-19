import React, { Component } from 'react';
import GameActions from '../actions/game';
import GameStore from '../stores/game';


class InitializeGame extends Component {
	constructor() {
		super();
		this.handleInitGame = this.handleInitGame.bind(this);
		this.onStatusChange = this.onStatusChange.bind(this);
	}

	onStatusChange(status) {
		if (status.game) {
			this.setState({ game: status.game });
			if (status.game.initialized) {
				window.location.href = '/play/' + status.game._id;
			}
		}
		if (status.error) {
			console.log(status.error);
			this.setState({ error: status.error });
		}
	}

	componentDidMount() {
		GameStore.listen(this.onStatusChange)
	}

	handleInitGame() {
		GameActions.initGame(this.props.game);
	}

	renderStart() {
		return (
			<div className="startGame">
				<h3>Have all your players joined? Then, it may be time to...</h3>
				<button className="btn initGame createBtn" onClick={this.handleInitGame}>Start Game</button>
			</div>
		);
	}

	renderWaiting() {
		let player = null;
		for (let i in this.props.game.players) {
			if (this.props.game.players[i]._id !== this.props.game.startedBy) {
				continue;
			}
			player = this.props.game.players[i];
		}
		return (
			<div className="startGame">
				<h3>Waiting for {player.name} to start the game.</h3>
			</div>
		);
	}

	render() {
		return (
			<div className="initGame">
				{this.props && this.props.user._id === this.props.game.startedBy ? this.renderStart() : this.renderWaiting()}
			</div>
		);
	}
}

export default InitializeGame;