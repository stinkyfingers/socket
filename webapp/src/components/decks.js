import React, { Component } from 'react';
import DeckStore from '../stores/deck';
import DeckActions from '../actions/deck';
import '../css/deck.css';

class List extends Component {

	constructor() {
		super()
		this.onStatusChange = this.onStatusChange.bind(this);
	}

	onStatusChange(status) {
		if (status.decks){
			this.setState({ decks: status.decks });
		}
	}

	componentWillMount() {
		DeckActions.allDecks();
	}

	componentDidMount() {
		DeckStore.listen(this.onStatusChange)
	}

	handleClick(id, action) {
		window.location.replace('/' + action + '/' + id);
	}

	renderDeckList() {
		let rows = [];
		for (let i in this.state.decks) {
			if (!this.state.decks) {
				continue;
			}
			rows.push(<tr key={'deck' + i}>
				<td>
					<button className='std play' onClick={this.handleClick.bind(null,this.state.decks[i].id, 'play')}>Play</button>
					<button className='std edit' onClick={this.handleClick.bind(null,this.state.decks[i].id, 'edit')}>Edit</button>
				</td>
				<td>{this.state.decks[i].deckType.name}</td>
				<td>{this.state.decks[i].deckType.description}</td>
				<td>{this.state.decks[i].user.username}</td>
			</tr>);
		}
		return (
			<table className='deckList'>
				<thead>
					<tr>
						<th></th>
						<th>Name</th>
						<th>Desc</th>
						<th>User</th>
					</tr>
				</thead>
				<tbody>
					{rows}
				</tbody>
			</table>
		);
	}


	render() {
		return (
			<div className="edit">
			{this.state && this.state.decks ? this.renderDeckList() : null}
			</div>
		);
	}
}

export default List;
