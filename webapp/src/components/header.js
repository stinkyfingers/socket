import React, { Component } from 'react';
import Login from './login';
import Logout from './logout';
import GameActions from '../actions/game';
import UserActions from '../actions/user';
import UserStore from '../stores/user';
import '../css/header.css';

class Header extends Component {
  constructor() {
    super()
    this.handleCancel = this.handleCancel.bind(this);
    this.statusUpdate = this.statusUpdate.bind(this);
  }

  statusUpdate(status) {
    if (status.user) {
      this.setState({ user: status.user });
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


  userDisplay() {
    return(
      <div className="userDisplay">
        Player: {this.state.user.name}
      </div>
    );
  }

  renderNav() {
    return (
      <div className='nav'>
        <ul className='navList'>
          <li><a href='/'>Find a Game</a></li>
          <li><a href='/create'>Start a Game</a></li>
          <li><a href='/decks'>Decks</a></li>
        </ul>
      </div>
    );
  }


  handleCancel() {
    GameActions.exitGame(this.props.game);
    window.location.href = '/';
  }

  renderCancelGame() {
    if (this.props.game.startedBy !== this.state.user._id) {
      return null;
    }
    return (<button className="cancel btn submit" onClick={this.handleCancel}>Cancel Game</button>);
  }

  render() {
    return (
      <div className="App">
        <div className="App-header">
          <h1>The Difference Between</h1>
          {this.renderNav()}
          {this.state && this.state.user ? this.userDisplay() : null }
          {this.state && this.state.user ? null : <Login className="login"  />}
          {this.state && this.state.user ? <Logout className="login" /> : null}
          {this.state && this.state.game ? this.renderCancelGame() : null}
        </div>
      </div>
    );
  }
}

export default Header;
