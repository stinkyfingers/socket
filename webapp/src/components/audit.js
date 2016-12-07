import React, { Component } from 'react';
import CardActions from '../actions/card';
import UserActions from '../actions/user';
import CardStore from '../stores/card';
import UserStore from '../stores/user';
import '../css/audit.css';
import classNames from 'classnames';


class AuditCards extends Component {
	constructor() {
		super();
		this.handleClick = this.handleClick.bind(this);
		this.statusUpdate = this.statusUpdate.bind(this);
	}

	statusUpdate(status) {
		if (status.cards) {
			this.setState({ cards: status.cards });
		}
		if (status.dealerCards) {
			this.setState({ dealerCards: status.dealerCards });
		}
		if (status.user) {
			this.setState({ user: status.user });
		}
	}

	componentWillMount() {
		CardActions.unreviewed()
		UserActions.getUser();
	}

	componentDidMount() {
    	this.unlisten = CardStore.listen(this.statusUpdate);
    	this.unlistenUser = UserStore.listen(this.statusUpdate);
	}

	componentWillUnmount() {
		this.unlisten();
		this.unlistenUser();
	}

	handleClick(val, type, approval, index) {
		val.corporateApproved = approval;
		CardActions.approve(val, type, this.state.user);
		let cards = this.state.cards;
		cards.splice(index ,1);
		this.setState({ cards });
	}

	renderCards() {
		const cards = [];
		for (const i in this.state.cards) {
			if (!this.state.cards[i]) {
				continue;
			}
			cards.push(<tr key={'card' + i} className={classNames('cardRow')}>
				<td>{this.state.cards[i].phrase}</td>
				<td className="approve" onClick={() => this.handleClick(this.state.cards[i], 'card', true, i)}>APPROVE</td>
				<td className="disapprove" onClick={() => this.handleClick(this.state.cards[i], 'card', false, i)}>DISAPPROVE</td>
				</tr>)
		}
		return (
			<table className="cardTable">
				<thead><tr><th colSpan="3">Player Cards</th></tr></thead>
				<tbody>
					{cards}
				</tbody>
			</table>
		);
	}

	renderDealerCards() {
		const cards = [];
		for (const i in this.state.dealerCards) {
			if (!this.state.dealerCards[i]) {
				continue;
			}
			cards.push(<tr key={'dealerCard' + i} className={classNames('cardRow')}>
				<td>{this.state.cards[i].phrase}</td>
				<td className="approve" onClick={() => this.handleClick(this.state.dealerCards[i], 'dealerCard', true, i)}>APPROVE</td>
				<td className="disapprove" onClick={() => this.handleClick(this.state.dealerCards[i], 'dealerCard', false, i)}>DISAPPROVE</td>
				</tr>)
		}
		return (
			<table className="cardTable">
				<thead><tr><th colSpan="3">Dealer Cards</th></tr></thead>
				<tbody>
					{cards}
				</tbody>
			</table>
		);
	}

	render() {
		return (
			<div className="cardsElement">
				<h2>Un-Audited Cards</h2>
				<div className="playerCardsDivider"></div> 
				{this.state && this.state.cards ? this.renderCards() : null}
				{this.state && this.state.cards ? this.renderDealerCards() : null}
			</div>
		);
	}
}

export default AuditCards;