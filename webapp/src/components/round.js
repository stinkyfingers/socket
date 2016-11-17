import React, { Component } from 'react';
// import GameActions from '../actions/game';
// import GameStore from '../stores/game';
import DealerCards from './dealerCards';
import Cards from './cards';


class Round extends Component {
	// constructor() {
		// super();
		// this.onStatusChange = this.onStatusChange.bind(this);
	// }

	// onStatusChange(status) {
	// 	console.log(status)
	// 	if (status.game) {
	// 		console.log(status)
	// 		this.setState({ game: status.game });
	// 	}
	// 	if (status.error) {
	// 		console.log(status.error)
	// 		// this.setState({ error: status.error });
	// 	}
	// }

	componentWillMount() {
	}

	componentDidMount() {
		// GameStore.listen(this.onStatusChange);
	}

	renderVotes() {
		let cards = [];
		for (const i in this.props.game.round.options) {
			if (!this.props.game.round.options) {
				continue;
			}
			cards.push(<div key={'choice' + i}className="optionCard card" onClick={this.handleClick} data-value={this.props.game.round.options[i].card.phrase}>{this.props.game.round.options[i].card.phrase}</div>)
		}
		return (
			<div className="options">
				<DealerCards dealerCards={this.props.game.round.dealerCards} />
				<h3>Options</h3>
				{cards}
			</div>
		);
	}

	renderRound() {
		if (!this.props.game.round.dealerCards) {
			// this.setState({ error: 'not enough dealer cards '});
			// return;
		}
		if (!this.props.game.round.cards) {
			// this.setState({ error: 'not enough player cards '});
			// return;
		}

		if (this.props.game.round.options && this.props.game.round.options.length > 0) {
			console.log("TIME TO VOTE");
			return this.renderVotes();
		} else {
			return this.renderCards();
		}
	}

	renderCards() {
		let cards = [];
		for (const i in this.props.game.players) {
			if (!this.props.game.players[i]) {
				continue;
			}
			if (this.props.game.players[i]._id === this.props.user._id) {
				cards = this.props.game.players[i].hand;
			}
		}

		return (
			<div className="play">
				<DealerCards dealerCards={this.props.game.round.dealerCards} />
				<Cards cards={cards} user={this.props.user} game={this.props.game} />
			</div>
		);

	}

	render() {
		return (
			<div className="play">Round: 
				{this.props && this.props.game && this.props.game.round ? this.renderRound() : null}
			</div>
		);
	}
}

export default Round;