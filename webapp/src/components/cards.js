import React, { Component } from 'react';
import GameActions from '../actions/game';
import '../css/card.css';

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
		if (this.props.cards.length !== 3) {
			console.log('not enough cards');
			// TODO handle error
			return
		}
		const cards = [];
		for (const i in this.props.cards) {
			if (!this.props.cards[i]) {
				continue;
			}
			if (this.playedCard && this.playedCard.phrase === this.props.cards[i].phrase){
				this.props.cards.splice(i);
			} else{
				cards.push(<div key={'card' + i} className="playerCard card" onClick={() => {this.handleClick(this.props.cards[i])}}>{this.props.cards[i].phrase}</div>)
			}
		}
		return (
			<div className="playerCardsContainer">
				{cards}
			</div>
		);
	}

	renderPlayedCard(){
		if (!this.playedCard) {
			return null;
		}
		return (
			<div className="playedCard card">
				{this.playedCard.phrase}
			</div>
		);
	}

	render() {
		return (
			<div className="cardsElement">
				<div className="playerCardsDivider"></div> 
				<h3>Click to play</h3>
				{this.props && this.props.cards ? this.renderCards() : null}
				{this.renderPlayedCard()}
			</div>
		);
	}
}

export default Cards;