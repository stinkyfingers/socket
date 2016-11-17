import React, { Component } from 'react';
// import GameActions from '../actions/game';
import GameStore from '../stores/game';
import '../css/card.css';

class DealerCards extends Component {
	constructor() {
		super();
		this.onStatusChange = this.onStatusChange.bind(this);
	}

	onStatusChange(status) {
		if (status.game) {
			this.setState({ game: status.game });
		}
		if (status.error) {
			console.log(status.error)
			this.setState({ error: status.error });
		}
	}

	componentWillMount() {
		this.setState({ cards: this.props.dealerCards });
	}

	componentDidMount() {
		GameStore.listen(this.onStatusChange);
	}

	renderCards() {
		if (this.state.cards.length !== 2) {
			console.log('not enough dealer cards');
			this.setState({ error: 'not enough dealer cards' });
		}
		const cards = (
			<div className="dealerCards">
				<span>The difference between</span>
				<div className="dealerCard card">{this.state.cards[0].phrase}</div>
				<span>and</span>
				<div className="dealerCard card">{this.state.cards[1].phrase}</div>
				<span>is...</span>
			</div>
			);
		return cards;
	}

	render() {
		return (
			<div className="play">Dealer Cards: 
				{this.state && this.state.cards ? this.renderCards() : null}
			</div>
		);
	}
}

export default DealerCards;