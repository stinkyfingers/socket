import React, { Component } from 'react';
import GameActions from '../actions/game';
import GameStore from '../stores/game';
import Initialize from './initializeGame';
import Round from './round';
import FinalScore from './finalScore';
import UserActions from '../actions/user';
import UserStore from '../stores/user';


class Play extends Component {
	constructor() {
		super();
		this.onStatusChange = this.onStatusChange.bind(this);
	}

	onStatusChange(status) {
		if (status.game) {
			this.setState({ game: status.game });
		}
		if (status.user) {
			this.setState({ user: status.user });
		}
	}

	componentWillMount() {
    	const u = location.href;
    	const index = u.lastIndexOf("/") + 1;
    	const id = u.substr(index);
    	GameActions.connect(id);
		UserActions.getUser();
	}

	componentDidMount() {
		this.unsubscribe = GameStore.listen(this.onStatusChange);
		this.unlisten = UserStore.listen(this.onStatusChange);
	}

	componentWillUnmount() {
		this.unsubscribe();
		this.unlisten();
	}

	renderGame() {
		if (this.state.game.initialized === false) {
			return (<Initialize game={this.state.game} user={this.state.user}/>);
		}
		return (<Round game={this.state.game} user={this.state.user} />);

	}

	render() {
		return (
			<div className="play">
				{this.state && this.state.game && this.state.game.finalScore ? <FinalScore game={this.state.game} user={this.state.user} /> : null}
				{this.state && this.state.game && !this.state.game.finalScore ? this.renderGame() : null}
			</div>
		);
	}
}

export default Play;