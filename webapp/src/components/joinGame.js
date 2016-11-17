import React, { Component } from 'react';
import GameActions from '../actions/game';
import GameStore from '../stores/game';

class JoinGame extends Component {
	constructor() {
		super();
		this.onStatusChange = this.onStatusChange.bind(this);
		this.handleClick = this.handleClick.bind(this);
	}

	onStatusChange(status) {
		if (status.error) {
			console.log(status.error)
			// this.setState({ error: status.error });
		}
	}

	componentWillMount() {
	}

	componentDidMount() {
		GameStore.listen(this.onStatusChange);
	}

	handleClick() {
		console.log('click', this.props.game)
		if (this.props.game.intialized) {
			console.log("Already started")
			// TODO - handle error
			return;
		}

		GameActions.joinGame(this.props.game, this.props.user);
	}


	render() {
		return (
			<div className="joinGame">Join Game {this.props.game._id}
				<button onClick={this.handleClick}>Join</button>
			</div>
		);
	}
}

export default JoinGame;