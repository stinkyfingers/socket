import React, { Component } from 'react';
import UserActions from '../actions/user';
import '../css/header.css'
import '../App.css'

class PasswordReset extends Component {
	constructor() {
		super();
		this.handleSubmit = this.handleSubmit.bind(this);
		this.handleField = this.handleField.bind(this);
	}

	handleSubmit() {
		UserActions.passwordReset(this._email);
	}

	handleField(e) {
		this._email = e.target.value;
	}

	handleShow(val) {
		this._reset = val;
		this.setState({ reset: val });
	}

	renderReset() {
		if (!this.state || !this.state.reset) {
			return null;
		}
		return (
			<div className="passwordResetForm">
				{this.state && this.state.error ? this.state.error.message : null}
				<label htmlFor="email">
					<input type="text" name="email" placeholder="Enter email..." onChange={this.handleField} />
				</label>
				<button className="btn submit" onClick={this.handleSubmit}>Send New Password</button>
				<button className="btn cancel" onClick={this.handleShow.bind(this, false)} >Cancel</button>
			</div>
		);
	}

	renderForgot() {
		if (this.state && this.state.reset) {
			return null;
		}
		return (<button className="btn show" onClick={this.handleShow.bind(this, true)}>Forgot Password?</button>);
	}

	render() {
		return (
			<div className="passwordReset">
				{this.renderReset()}
				{this.renderForgot()}
			</div>
		);
	}
}

export default PasswordReset;