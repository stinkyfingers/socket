import React, { Component } from 'react';
import GameActions from '../actions/game';
import '../css/card.css';
import classNames from 'classnames';


class Cards extends Component {
	constructor() {
		super();
		this.handleClick = this.handleClick.bind(this);
	}

	handleClick(card) {
		const play = {
			playType: 'play',
			player: this.props.user,
			card: card
		}
		this.playedCard = play.card;
		GameActions.play(this.props.game._id, this.props.user, play);
	}

	renderCards() {
		if (this.props.cards.length < 3) {
			console.log('not enough cards');
			// TODO handle error
			return
		}
		const cards = [];
		for (const i in this.props.cards) {
			if (!this.props.cards[i]) {
				continue;
			}
			cards.push(<div key={'card' + i} className={classNames('playerCard', 'card', this.playedCard && this.playedCard.phrase === this.props.cards[i].phrase ? 'active' : null)} onClick={() => {this.handleClick(this.props.cards[i])}}>{this.props.cards[i].phrase}</div>)
		}
		return (
			<div className="playerCardsContainer">
				{cards}
			</div>
		);
	}

	render() {
		return (
			<div className="cardsElement">
				<div className="playerCardsDivider"></div> 
				<h3>Click to play</h3>
				{this.props && this.props.cards ? this.renderCards() : null}

			</div>
		);
	}
}

export default Cards;