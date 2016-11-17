import React, { Component } from 'react';
import DeckStore from '../stores/deck';
import DeckActions from '../actions/deck';
import '../css/deck.css';

class Edit extends Component {

	constructor() {
		super()
		this.onStatusChange = this.onStatusChange.bind(this);
		this.handleChange = this.handleChange.bind(this);
		this.handleEdit = this.handleEdit.bind(this);
		this.handleSaveDeck = this.handleSaveDeck.bind(this);
		this.handleChangeCard = this.handleChangeCard.bind(this);
	}

	onStatusChange(status) {
		if (status.deck) {
			this.setState({ deck: status.deck });
		}
	}

	componentWillMount() {
		const id = window.location.pathname.replace('/edit/','');
		DeckActions.getDeck(id);
	}

	componentDidMount() {
		DeckStore.listen(this.onStatusChange)
	}

	handleChange(e) {
		let deck = this.state.deck;
		deck[e.target.getAttribute('name')] = e.target.value;
		this.setState({ deck: deck });
	}

	handleChangeCard(e) {
		let deck = this.state.deck;
		for (let i in deck.positions) {
			if (!deck.positions || this.state.position.name !== deck.positions[i].name) {
				continue;
			}
			for (let j in deck.positions[i].cards) {
				if (!deck.positions[i].cards[j]) {
					continue;
				}
				if (j === e.target.getAttribute('data-key')) {
					deck.positions[i].cards[j].phrase = e.target.value;
				}
			}
		}
		this.setState({ deck: deck})
	}

	handleEdit(position) {
		this.setState({ cards: position.cards, position: position });
	}

	handleSaveDeck() {
		DeckActions.saveDeck(this.state.deck);
	}

	renderEditor() {
		const deck = this.state.deck;
		let positions = [];
		for (let i in deck.positions) {
			if (!deck.positions[i]) {
				continue;
			}
			positions.push(
				<div key={'position' + i} >
					<label htmlFor={'positions[' + i + ']'}>Cards Name:
						<input className='editForm' type='text' name={'positions[' + i + ']'} defaultValue={deck.positions[i].name} onChange={this.handleChange} />
						<button onClick={this.handleEdit.bind(null, deck.positions[i])} className='std edit'>Edit Cards</button>
					</label>
				</div>
			);
		}
		return(
			<div className='edit'>
				<label htmlFor='name'>Deck Name:
					<input className='editForm' type='text' name='name' defaultValue={deck.deckType.name} onChange={this.handleChange} />
				</label>
				<label htmlFor='description'>Deck Description:
					<input className='editForm' type='text' name='description' defaultValue={deck.deckType.description} onChange={this.handleChange} />
				</label>
				{positions}
				<button className='std save' onClick={this.handleSaveDeck}>Save</button>
			</div>
		);
	}

	renderCards() {
		let cards = [];
		for (let i in this.state.cards) {
			if (!this.state.cards) {
				continue;
			}
			cards.push(<input key={'card' + i } value={this.state.cards[i].phrase} data-key={i} onChange={this.handleChangeCard} />);
		}
		return (
			<div className='cards'>{cards}</div>
		);
	}


	render() {
		return (
			<div className="edit">
				{this.state && this.state.deck ? this.renderEditor() : null}
				{this.state && this.state.cards ? this.renderCards() : null}
			</div>
		);
	}
}

export default Edit;
