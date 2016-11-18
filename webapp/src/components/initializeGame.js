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

	render() {
		return (
			<div className="initGame">
				<h3>Have all your players joined? Then, it may be time to...</h3>
				<button className="initGame" onClick={this.handleInitGame}>Start Game
				</button>
			</div>
		);
	}
}

export default InitializeGame;