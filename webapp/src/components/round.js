import React, { Component } from 'react';
import GameActions from '../actions/game';
// import GameStore from '../stores/game';
import DealerCards from './dealerCards';
import Cards from './cards';
import FinalScore from './finalScore';
import classNames from 'classnames';


class Round extends Component {
	constructor() {
		super();
		this.handleClick = this.handleClick.bind(this);
	}


	handleClick(e) {
		const play = {
			playType: 'vote',
			card: e,
			player: this.props.user
		}
		GameActions.play(this.props.game._id, this.props.user, play);
		this.active = e.phrase
	}


	renderVotes() {
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

			cards.push(<div key={'choice' + i} className={classNames('optionCard','card', this.props.game.round.options[i].card.phrase === this.active ? 'active' : null)} onClick={() => this.handleClick(this.props.game.round.options[i].card)} data-value={this.props.game.round.options[i].card}>
				{this.props.game.round.options[i].card.phrase} {yours}
			</div>);
		}
		return (
			<div className="optionsContainer">
				<DealerCards dealerCards={this.props.game.round.dealerCards} />
				<h3>Options (Click to vote on your favorite)</h3>
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

	renderPreviousResults() {
		if (!this.props.game.round.mostRecentResults || !this.props.game.round.mostRecentResults.dealerCards || this.props.game.round.mostRecentResults.dealerCards.length === 0) {
			return null
		}
		const p = this.props.game.round.mostRecentResults;
		const votes = {};
		let tally = {};
		for (let id in p.votes) {
			if (!p.votes[id]) {
				continue;
			}
			if (!tally.id) {
				tally.id = 1;
			} else {
				tally.id++;
			}
			votes.id = (<div className="previousResult" key={'votes' + id}>
					The difference between<span className="previousCard">{p.dealerCards[0].phrase}</span>
					 and<span className="previousCard">{p.dealerCards[1].phrase}</span> is 
				<span className="previousCard">{p.votes[id].card.phrase}.</span>
				(Played by {p.votes[id].card.player.name}). <span className="total">Total: {tally.id}</span></div>);
		}

		const out = [];
		for (let o in votes) {
			if (!votes[o]) {
				continue;
			}
			out.push(votes[o]);
		}

		return (	
			<div className="previous">
				<h3>Previous Round Results</h3>
				{out}
			</div>
		);
	}

	render() {
		return (
			<div className="play">
				{this.props && this.props.game && this.props.game.round && !this.props.game.finalScore ? this.renderRound() : null}
				{this.props && this.props.game && this.props.game.finalScore ? <FinalScore game={this.props.game} user={this.props.user} /> : null}
				<div className="playerCardsDivider"></div>
				{this.props && this.props.game && this.props.game.round && this.props.game.round.mostRecentResults ? this.renderPreviousResults() : null}
			</div>
		);
	}
}

export default Round;