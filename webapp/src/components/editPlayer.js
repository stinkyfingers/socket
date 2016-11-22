import React, { Component } from 'react';
import UserActions from '../actions/user';
import UserStore from '../stores/user';
import '../css/user.css';

class EditPlayer extends Component {

	constructor() {
		super();

		this.handleChange = this.handleChange.bind(this);
		this.handleClick = this.handleClick.bind(this);
		this.statusUpdate = this.statusUpdate.bind(this);
	}

	statusUpdate(status) {
		if (status.user) {
			this.setState({ user: status.user });
		} 
		if (status.user === null) {
			this.setState({ user: {} });
			return;
		}
}

	componentWillMount() {
		UserActions.getUser();
	}

	componentDidMount() {
		this.unlisten = UserStore.listen(this.statusUpdate);
	}

	componentWillUnmount() {
		this.unlisten();
	}


	handleChange(e) {
		const type = e.target.name;
		const val = e.target.value;

		let user = this.state.user;
		user[type] = val;
		this.setState({ user: user });
	}

	handleClick() {
		UserActions.saveUser(this.state.user);
	}

	renderEdit() {
		return (
			<div className="editContainer">
				<label htmlFor="name">Name: 
					<input type="text" onChange={this.handleChange} name="name" defaultValue={this.state.user.name} />
				</label>
				<label htmlFor="email">Email: 
					<input type="text" onChange={this.handleChange} name="email" defaultValue={this.state.user.email} />
				</label>
				<label htmlFor="password">Password {this.state && this.state.user && this.state.user._id ? '(leave blank to not change)' : ''}: 
					<input type="text" onChange={this.handleChange} name="password" defaultValue={this.state.user.password} />
				</label>
				<button className="btn submit" onClick={this.handleClick}>Save</button>
			</div>
		);
	}



	render() {
		return (
			<div className="editPlayer">
				{this.state ? this.renderEdit() : null}
			</div>
		);
	}
}

export default EditPlayer;