import React, { Component } from 'react';
import UserActions from '../actions/user';
import UserStore from '../stores/user';
import '../css/header.css'
import '../App.css'

class Logout extends Component {
	constructor() {
		super();
		this.onStatusChange = this.onStatusChange.bind(this);
		this.handleLogout = this.handleLogout.bind(this);
	}

	onStatusChange(status) {
		if (status.user) {
			this.setState({ user: status.user });
		}
		if (status.error) {
			this.setState({ error: status.error });
		}
	}

	componentDidMount() {
		UserStore.listen(this.onStatusChange);
	}

	handleLogout() {
		UserActions.unsetUser();
	}

	render() {
		return (
			<div className="login">
				<button className="btn submit" onClick={this.handleLogout}>Logout</button>
			</div>
		);
	}
}

export default Logout;