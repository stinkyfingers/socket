import React, { Component } from 'react';
import GameActions from '../actions/game';
import GameStore from '../stores/game';
import UserStore from '../stores/user';
import JoinGame from './joinGame';

class FindGame extends Component {
	constructor() {
		super();
		this.onStatusChange = this.onStatusChange.bind(this);
		this.handleClick = this.handleClick.bind(this);
		this.handleChange = this.handleChange.bind(this);
	}

	onStatusChange(status) {
		if (status.game) {
			this.setState({ game: status.game });
		}
		if (status.error) {
			console.log(status.error)
			this.setState({ error: status.error });
		}
		if (status.user) {
			this.setState({ user: status.user });
		}
	}

	componentWillMount() {
	}

	componentDidMount() {
		GameStore.listen(this.onStatusChange);
		UserStore.listen(this.onStatusChange);
	}

	handleClick() {
		GameActions.getGame(this.state.id);
	}

	handleChange(e) {
		const id = e.target.value;
		this.setState({ id: id });
	}

	renderPlayGame() {
		let alreadyJoined = false;
		for (let i in this.state.game.players) {
			if (!this.state.game.players[i]) {
				continue;
			}
			if (this.state.game.players[i]._id === this.state.user._id) {
				alreadyJoined = true;
			}
		}
		if (!alreadyJoined) {
			return (<JoinGame user={this.state.user} game={this.state.game} />);
		}
		return (
			<div>
				<a href={'/create/' + this.state.game._id} >Play</a>
			</div>
		);
	}

	render() {
		return (
			<div className="findGame">Find Game
				<label htmlFor="id">
					<input name="id" onChange={this.handleChange} />
				</label>
				<button onClick={this.handleClick}>Search</button>
				{this.state && this.state.game && this.state.game.initialized === false ? this.renderPlayGame() : null}
				{this.state && this.state.game && this.state.game.initialized === true ? <div className="started">Game has already started</div> : null}
			</div>
		);
	}
}

export default FindGame;