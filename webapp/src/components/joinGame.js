import React, { Component } from 'react';
import GameActions from '../actions/game';
import GameStore from '../stores/game';

class JoinGame extends Component {
	constructor() {
		super();
		this.onStatusChange = this.onStatusChange.bind(this);
		this.handleClick = this.handleClick.bind(this);
	}

	componentDidMount() {
		GameStore.listen(this.onStatusChange);
	}

	handleClick() {
		if (this.props.game.intialized) {
			this.setState({ error: {message: "Game already started", error: "Start a new game"}});
			return;
		}

		GameActions.joinGame(this.props.game, this.props.user);
	}


	render() {
		return (
			<div className="joinGame">
				<h3>Game {this.props.game._id} found!</h3>
				<button className="joinGame btn" onClick={this.handleClick}>Join Game</button>
			</div>
		);
	}
}

export default JoinGame;