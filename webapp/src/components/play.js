import React, { Component } from 'react';
import GameActions from '../actions/game';
import GameStore from '../stores/game';
import Initialize from './initializeGame';
import Round from './round';

class Play extends Component {
	constructor() {
		super();
		this.onStatusChange = this.onStatusChange.bind(this);
	}

	onStatusChange(status) {
		if (status.game) {
			// console.log('updating game', status.game)
			this.setState({ game: status.game });
		}
		if (status.error) {
			console.log(status.error)
			// this.setState({ error: status.error });
		}
	}

	componentWillMount() {
    	const u = location.href;
    	const index = u.lastIndexOf("/") + 1;
    	const id = u.substr(index);
    	GameActions.connect(id);
	}

	componentDidMount() {
		this.unsubscribe = GameStore.listen(this.onStatusChange);
	}

	componentWillUnmount() {
		this.unsubscribe();
	}

	renderGame() {
		if (this.props.game.initialized === false) {
			return (<Initialize game={this.props.game} />);
		}
		return (<Round game={this.props.game} user={this.props.user} />);

	}

	render() {
		// console.log(this.props, this.state)
		return (
			<div className="play">Play: 
			{this.props && this.props.game ? this.renderGame() : null}
			</div>
		);
	}
}

export default Play;