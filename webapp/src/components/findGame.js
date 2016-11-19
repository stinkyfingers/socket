import React, { Component } from 'react';
import GameActions from '../actions/game';
import GameStore from '../stores/game';
import UserStore from '../stores/user';
import JoinGame from './joinGame';
import '../css/findGame.css'

class FindGame extends Component {
	constructor() {
		super();
		this.onStatusChange = this.onStatusChange.bind(this);
		this.handleClick = this.handleClick.bind(this);
		this.handleChange = this.handleChange.bind(this);
		this.playerJoined = false;
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

		if (this.state && this.state.game && this.state.user) {
			for (const i in this.state.game.players) {
				if (this.state.user._id === this.state.game.players[i]._id) {
					this.playerJoined = true;
				}
			}
		}
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
				<h3>You're enrolled in game {this.state.game._id}. Click 
				<span> <a href={'/play/' + this.state.game._id} >here</a> </span>
				to begin.</h3>
			</div>
		);
	}

	render() {
		if (this.state && this.state.game && this.state.user) {
			for (const i in this.state.game.players) {
				if (this.state.user._id === this.state.game.players[i]._id) {
					this.playerJoined = true;
				}
			}
		}
		return (
			<div className="findGame">Find Game:
				<label htmlFor="id">
					<input name="id" placeholder="Enter ID number..." onChange={this.handleChange} />
				</label>
				<button className="btn findGame" onClick={this.handleClick}>Search</button>
				{this.state && this.state.game && this.state.game.initialized === false ? this.renderPlayGame() : null}
				{this.state && this.state.game && this.state.game.initialized && !this.playerJoined === true ? <div className="started">Game has already started</div> : null}
				{this.state && this.state.game && this.state.game.initialized && this.playerJoined === true ? this.renderPlayGame() : null}
			</div>
		);
	}
}

export default FindGame;