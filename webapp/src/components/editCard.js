import React, { Component } from 'react';
import CardActions from '../actions/card';
import CardStore from '../stores/card';
import UserActions from '../actions/user';
import UserStore from '../stores/user';
import '../css/card.css';

class EditCard extends Component {

	constructor() {
		super();

		this.handleChange = this.handleChange.bind(this);
		this.handleClick = this.handleClick.bind(this);
		this.statusUpdate = this.statusUpdate.bind(this);
		this.handleType = this.handleType.bind(this);
		this.saved = false;
	}

	statusUpdate(status) {
		if (status.card) {
			this.setState({ card: status.card });
		}
		if (status.user) {
			this.setState({ user: status.user });
		}
	}

	componentWillMount() {
		UserActions.getUser();
		this.setState({ type: 'Card' });
	}

	componentDidMount() {
		this.unlisten = CardStore.listen(this.statusUpdate);
		this.userUnlisten = UserStore.listen(this.statusUpdate);
	}

	componentWillUnmount() {
		this.unlisten();
		this.userUnlisten();
	}


	handleChange(e) {
		this.saved = false;
		const type = e.target.name;
		const val = e.target.value;

		let card = this.state && this.state.card ? this.state.card : {createdBy: this.state.user._id};
		card[type] = val;
		this.setState({ card: card });
	}

	handleClick() {
		CardActions.create(this.state.card, this.state.type);
		this.saved = true;
	}

	handleType() {
		const type = this.state.type === 'card' ? 'Dealer' : 'Card';
		this.setState({ type });
		this.saved = false;
	}

	renderEdit() {
		return (
			<div className="editContainer">
				<label htmlFor="phrase">Phrase: 
					<input type="text" onChange={this.handleChange} name="phrase" />
				</label>

				<div className="type">Card Type: {this.state.type}</div>
				<label className="switch">
					<input type="checkbox" onChange={this.handleType} />
					<div className="slider"></div>
				</label>
				<button className="btn submit" onClick={this.handleClick}>Save</button>
			</div>
		);
	}

	renderSuccess() {
		return (
			<div className="success">
				Successfully saved!
			</div>
		);
	}


	render() {
		return (
			<div className="editPlayer">
				{this.state ? this.renderEdit() : null}
				{this.state && this.state.card && this.saved ? this.renderSuccess() : null}
			</div>
		);
	}
}

export default EditCard;