import React, { Component } from 'react';
import GameStore from '../stores/game';
import UserStore from '../stores/user';
import '../App.css'

class Error extends Component {
  constructor() {
    super()
    this.statusUpdate = this.statusUpdate.bind(this);
  }

  statusUpdate(status) {
    if (status.error) {
      this.setState({ error: status.error });
    }
  }

  componentDidMount() {
    this.unlisten = UserStore.listen(this.statusUpdate);
    this.unlisten = GameStore.listen(this.statusUpdate);
  }

  componentWillUnmount() {
    this.unlisten();
  }

  renderError() {
    return (
      <div className="errorContainer">
        {this.state.error.message}: {this.state.error.error}
      </div>);
  }

  render() {
    return (
      <div className="error">
        {this.state && this.state.error ? this.renderError() : null}
      </div>
    );
  }
}

export default Error;
