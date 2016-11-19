import React, { Component } from 'react';
import GameActions from '../actions/game';
import GameStore from '../stores/game';
import InitializeGame from '../components/initializeGame';
import '../css/createGame.css';

class CreateGame extends Component {
	constructor() {
		super();
		this.handleNewGame = this.handleNewGame.bind(this);
		this.onStatusChange = this.onStatusChange.bind(this);
		this.handleRefresh = this.handleRefresh.bind(this);
	}

	onStatusChange(status) {
		if (status.game && !status.game.initialized) {
			this.setState({ game: status.game });
		}
		if (status.error) {
			console.log(status.error);
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
	}

	componentDidMount() {
    	GameStore.listen(this.onStatusChange);
	}

	handleNewGame(){
		GameActions.createGame(this.props.user);
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
		console.log(this.props)
		if (this.props.user._id === this.state.game.startedBy) {
			return (<InitializeGame game={this.state.game} user={this.props.user}/>);
		}
		return (<div className="waitTostart">Waiting to start...</div>);
	}

	render() {
		console.log(this.state)
		return (
			<div className="createGame">
				{this.state && this.state.game ? null : <button className="btn createBtn" onClick={this.handleNewGame}>Create New Game</button>}
				{this.state && this.state.game ? this.handlePlayerList() : null}
				{this.state && this.state.game && !this.state.game.initialized ? this.handleInitGameButton() : null}
			</div>
		);
	}
}

export default CreateGame;