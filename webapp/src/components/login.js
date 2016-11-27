import React, { Component } from 'react';
import UserActions from '../actions/user';
import UserStore from '../stores/user';
import '../css/header.css'
import '../App.css'

class Login extends Component {
	constructor() {
		super();
		this.onStatusChange = this.onStatusChange.bind(this);
		this.handleLogin = this.handleLogin.bind(this);
		this.handleField = this.handleField.bind(this);
	}

	onStatusChange(status) {
		if (status.user) {
			this.setState({ user: status.user });
		}
	}

	componentDidMount() {
		if (this.props.user === null) { //TODO /??? 
			return;
		}
		UserStore.listen(this.onStatusChange);
	}

	handleLogin() {
		UserActions.authenticate(this.state.user);
	}

	handleField(e) {
		const value = e.target.value;
		const field = e.target.name;
		const user = this.state && this.state.user ? this.state.user : {};
		user[field] = value;
		this.setState({ user });
	}

	renderLoginButton() {
		if (!this.state || !this.state.user || !this.state.user.name || !this.state.user.password) {
			return null;
		}
		return (
			<button className="btn submit" onClick={this.handleLogin}>Login</button>
		);
	}

	render() {
		return (
			<div className="login">
				{this.state && this.state.error ? this.state.error.message : null}
				<label htmlFor="name">Username:
					<input type="text" name="name" onChange={this.handleField} />
				</label>
				<label htmlFor="password">Password:
					<input type="password" name="password" onChange={this.handleField} />
				</label>
				{this.renderLoginButton()}
			</div>
		);
	}
}

export default Login;