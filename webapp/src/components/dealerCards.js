import React, { Component } from 'react';
import '../css/card.css';
import '../css/round.css';

class DealerCards extends Component {
	
	renderCards() {
		if (this.props.dealerCards.length !== 2) {
			console.log('not enough dealer cards');
			this.setState({ error: 'not enough dealer cards' });
		}
		const cards = (
			<div className="dealerCardsContainer">
				<span>The difference between</span>
				<div className="dealerCard card">{this.props.dealerCards[0].phrase}</div>
				<span>and</span>
				<div className="dealerCard card">{this.props.dealerCards[1].phrase}</div>
				<span>is...</span>
			</div>
			);
		return cards;
	}

	render() {
		return (
			<div className="dealerCardsElement">
				{this.props && this.props.dealerCards ? this.renderCards() : null}
			</div>
		);
	}
}

export default DealerCards;