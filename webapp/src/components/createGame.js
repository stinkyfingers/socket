import React, { Component } from 'react';
import GameActions from '../actions/game';
import GameStore from '../stores/game';
import UserActions from '../actions/user';
import UserStore from '../stores/user';
import InitializeGame from '../components/initializeGame';
import '../css/createGame.css';

class CreateGame extends Component {
	constructor() {
		super();
		this.handleNewGame = this.handleNewGame.bind(this);
		this.onStatusChange = this.onStatusChange.bind(this);
		this.handleRefresh = this.handleRefresh.bind(this);
		this.handleNumRounds = this.handleNumRounds.bind(this);
	}

	onStatusChange(status) {
		if (status.game && !status.game.initialized) {
			this.setState({ game: status.game });
		}
		if (status.error) {
			console.log(status.error);
		}
		if (status.user) {
			this.setState({ user: status.user });
		}
	}
	componentWillMount() {
		GameActions.getGameFromStorage();

		const u = location.href;
    	const index = u.lastIndexOf("/") + 1;
    	const id = u.substr(index);

    	if (id && id !== 'create' && this.state && id !== this.state.game._id) {
    		GameActions.getGame(id);
    	}
		UserActions.getUser();

	}

	componentDidMount() {
    	GameStore.listen(this.onStatusChange);
		this.unlisten = UserStore.listen(this.onStatusChange);
	}

	componentWillUnmount() {
		this.unlisten();
	}

	handleNewGame(){
		GameActions.createGame(this.state.user);
	}

	handleRefresh() {
		GameActions.getGame(this.state.game._id);
	}

	handlePlayerList(){
		const players = [];
		for (let i in this.state.game.players) {
			if (!this.state.game.players[i]) {
				continue;
			}
			players.push(<div key={'player' + i} className="player">{this.state.game.players[i].name}</div>);
		}

		return (
			<div className="players">
				<h3>Players for game <span className="id">{this.state.game._id}:</span></h3>
				{players}
				<button className="btn refresh createBtn" onClick={this.handleRefresh}>Refresh Player List</button>
			</div>
		);
	}

	handleInitGameButton() {
		if (this.state.user._id === this.state.game.startedBy) {
			return (<InitializeGame game={this.state.game} user={this.state.user}/>);
		}
		return (<div className="waitTostart">Waiting to start...</div>);
	}

	handleNumRounds(e) {
		let g = this.state.game;
		g.roundsInGame = parseInt(e.target.value, 10);
		GameActions.updateGame(g);
	}

	renderNumRounds() {
		let options = [];
		for (let i = 1; i < 11; i++) {
			options.push(<option key={'numRound' + i}>{i}</option>);
		}
		return(
			<div className="numRounds">
				<label>Number of Rounds:</label>
				<select onChange={this.handleNumRounds}>
				{options}
				</select>
			</div>
		);
	}

	render() {
		return (
			<div className="createGame">
				{this.state && this.state.game ? null : <button className="btn createBtn" onClick={this.handleNewGame}>Create New Game</button>}
				{this.state && this.state.game ? this.handlePlayerList() : null}
				{this.state && this.state.game ? this.renderNumRounds() : null}
				{this.state && this.state.game && !this.state.game.initialized ? this.handleInitGameButton() : null}
			</div>
		);
	}
}

export default CreateGame;