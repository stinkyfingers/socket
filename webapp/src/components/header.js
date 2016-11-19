import React, { Component } from 'react';
import Login from './login';
import Logout from './logout';
import GameActions from '../actions/game';
import '../css/header.css';

class Header extends Component {
  constructor() {
    super()
    this.handleCancel = this.handleCancel.bind(this);
  }

  userDisplay() {
    return(
      <div className="userDisplay">
        Player: {this.props.user.name}
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
    if (this.props.game.startedBy !== this.props.user._id) {
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
          {this.props && this.props.user ? this.userDisplay() : null }
          {this.props && this.props.user ? null : <Login className="login"  />}
          {this.props && this.props.user ? <Logout className="login" /> : null}
          {this.props && this.props.game ? this.renderCancelGame() : null}
        </div>
      </div>
    );
  }
}

export default Header;
