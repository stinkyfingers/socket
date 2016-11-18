import React, { Component } from 'react';
import GameActions from '../actions/game';
// import GameStore from '../stores/game';
import DealerCards from './dealerCards';
import Cards from './cards';
import FinalScore from './finalScore';


class Round extends Component {
	constructor() {
		super();
		// this.onStatusChange = this.onStatusChange.bind(this);
		this.handleClick = this.handleClick.bind(this);
	}

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

	handleClick(e) {
		// const card = e.target.dataset.value;
		const play = {
			playType: 'vote',
			card: e,
			player: this.props.user
		}
		GameActions.play(this.props.game._id, this.props.user, play);
	}

	renderVotes() {
		console.log(this.props.game)
		let cards = [];
		for (const i in this.props.game.round.options) {
			if (!this.props.game.round.options) {
				continue;
			}
			// mark player's card
			let yours = null;
			if (this.props.game.round.options[i].player._id === this.props.user._id) {
				yours = <span className="youPlayed">(You played this, genius)</span>;
			}

			cards.push(<div key={'choice' + i} className="optionCard card" onClick={() => this.handleClick(this.props.game.round.options[i].card)} data-value={this.props.game.round.options[i].card}>
				{this.props.game.round.options[i].card.phrase} {yours}
			</div>);
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

	renderPreviousRound() {
		const p = this.props.game.round.previous;
		console.log(p)
		return (<div className="previous">
			TODO - previous round - wins
			</div>
		);
	}

	render() {
		console.log(this.props.game)
		return (
			<div className="play">Round: 
				{this.props && this.props.game && this.props.game.round && !this.props.game.finalScore ? this.renderRound() : null}
				{this.props && this.props.game && this.props.game.finalScore ? <FinalScore game={this.props.game} user={this.props.user} /> : null}
				{this.props && this.props.game && this.props.game.round && this.props.game.round.previous ? this.renderPreviousRound() : null}
			</div>
		);
	}
}

export default Round;