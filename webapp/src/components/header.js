import React, { Component } from 'react';
import Login from './login';
import Logout from './logout';
import PasswordReset from './passwordReset';
import GameActions from '../actions/game';
import UserActions from '../actions/user';
import UserStore from '../stores/user';
import GameStore from '../stores/game';
import '../css/header.css';

class Header extends Component {
  constructor() {
    super()
    this.handleCancel = this.handleCancel.bind(this);
    this.statusUpdate = this.statusUpdate.bind(this);
    this.handleEditCard = this.handleEditCard.bind(this);
    this.handleReset = this.handleReset.bind(this);
  }

  statusUpdate(status) {
    if (status.user) {
      this.setState({ user: status.user });
    }
    if (status.game) {
      this.setState({ game: status.game });
    }
  }

  componentWillMount() {
    UserActions.getUser();
    GameActions.getGameFromStorage();
  }

  componentDidMount() {
    this.unlisten = UserStore.listen(this.statusUpdate);
    this.unlistenGame = GameStore.listen(this.statusUpdate);
  }

  componentWillUnmount() {
    this.unlisten();
    this.unlistenGame();
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
          <li><a href='/player'>{this.state && this.state.user ? 'Edit User' : 'Create User'} </a></li>
        </ul>
      </div>
    );
  }


  handleCancel() {
    GameActions.exitGame(this.props.game);
    window.location.href = '/';
  }

  handleEditCard() {
    window.location.href = '/card';
  }

  handleAuditCards() {
    window.location.href = '/audit';
  }

  handleReset() {
    GameActions.unsetGame();
    window.location.href = '/';
  }

  renderCancelGame() {
    if (this.state && this.state.user && this.state.game.startedBy !== this.state.user._id) {
      return null;
    }
    return (<button className="cancel btn submit" onClick={this.handleCancel}>Cancel Game</button>);
  }

  renderCreateCard() {
    return (<button className="cancel btn create" onClick={this.handleEditCard}>Create Card</button>);
  }

  renderAuditCards() {
    return (<button className="cancel btn create" onClick={this.handleAuditCards}>Audit Cards</button>);
  }

  render() {
    return (
      <div className="App">
        <div className="App-header">
          <img className="logo" src={require("../../public/logo.png")} alt="The Difference Between"/>
          {this.renderNav()}
          {this.state && this.state.user ? this.userDisplay() : null }
          {this.state && this.state.user ? null : <Login className="login" user={null} />}
          {this.state && this.state.user ? null : <PasswordReset className="reset" />}
          {this.state && this.state.user ? <Logout className="login" /> : null}
          {this.state && this.state.game ? this.renderCancelGame() : null}
          {this.state && this.state.user ? this.renderCreateCard() : null}
          {this.state && this.state.user && this.state.user.grandPoobah ? this.renderAuditCards(): null}
          <span className="wonky" onClick={this.handleReset}>Something's WoNkY</span>
        </div>
      </div>
    );
  }
}

export default Header;
