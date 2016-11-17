import React, { Component } from 'react';
import UserActions from '../actions/user';
import UserStore from '../stores/user';

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
		if (status.error) {
			this.setState({ error: status.error });
		}
	}

	componentDidMount() {
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
			<button className="button submit" onClick={this.handleLogin}>Login</button>
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