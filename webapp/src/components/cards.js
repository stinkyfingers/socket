import React, { Component } from 'react';
import GameActions from '../actions/game';
// import GameStore from '../stores/game';
import '../css/card.css';

class Cards extends Component {
	constructor() {
		super();
		// this.onStatusChange = this.onStatusChange.bind(this);
		this.handleClick = this.handleClick.bind(this);
	}

	// onStatusChange(status) {
		// if (status.game) {
		// 	this.setState({ game: status.game });
		// }
		// if (status.error) {
		// 	console.log(status.error)
		// 	this.setState({ error: status.error });
		// }
	// }

	componentWillMount() {
	}

	componentDidMount() {
		// GameStore.listen(this.onStatusChange);
	}

	handleClick(e) {
		const phrase = e.target.dataset.value;
		const play = {
			playType: 'play',
			player: this.props.user,
			card: {
				phrase: phrase
			}
		}
		GameActions.play(this.props.game._id, this.props.user, play);
	}

	renderCards() {
		if (this.props.cards.length !== 3) {
			console.log('not enough cards');
			// TODO handle error
			return
		}
		const cards = (
			<div className="playerCards">
				<div className="playerCard card" onClick={this.handleClick} data-value={this.props.cards[0].phrase}>{this.props.cards[0].phrase}</div>
				<div className="playerCard card" onClick={this.handleClick} data-value={this.props.cards[1].phrase}>{this.props.cards[1].phrase}</div>
				<div className="playerCard card" onClick={this.handleClick} data-value={this.props.cards[2].phrase}>{this.props.cards[2].phrase}</div>
			</div>
			);
		return cards;
	}

	render() {
		return (
			<div className="play">Player Cards: 
				{this.props && this.props.cards ? this.renderCards() : null}
			</div>
		);
	}
}

export default Cards;