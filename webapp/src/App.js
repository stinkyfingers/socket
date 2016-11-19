import React, { Component } from 'react';
import './App.css';
import UserStore from './stores/user';
import GameStore from './stores/game';
import UserActions from'./actions/user';
import Play from './components/play';
import Decks from './components/decks';
import Edit from './components/edit';
import CreateGame from './components/createGame';
import InitializeGame from './components/initializeGame';
import FindGame from './components/findGame';
import Header from './components/header';

class App extends Component {
  constructor() {
    super()
    this.onStatusChange = this.onStatusChange.bind(this);
    this.gameStatusChange = this.gameStatusChange.bind(this);
  }

  onStatusChange(status) {
    if (status.user) {
      this.setState({ user: status.user });
    }
    if (status.user === undefined) {
      this.setState({ user: undefined });
    }
  }

  gameStatusChange(status) {

    if (status.game && this._isMounted) {
      this.setState({ game: status.game });
    }

    // From game.play - avoid setState issue
    if (this._isMounted && status.gameUpdate) {
      this.setState({ game: status.gameUpdate });
    }
  }

  componentWillMount() {
    UserActions.getUser();
  }

  componentDidMount() {
    UserStore.listen(this.onStatusChange);
    this.unmountgame = GameStore.listen(this.gameStatusChange);
    this._isMounted = true;
  }

  componentWillUnmount() {
    this.unmountgame();
    this._isMounted = false;
  }

  renderNav() {
    return (
      <div className='nav'>
        <ul className='navList'>
          <li><a href='/'>Home</a></li>
          <li><a href='/decks'>Decks</a></li>
          <li><a href='/create'>Start New Game</a></li>
        </ul>
      </div>
    );
  }

  getRoute() {
    let Child;
    const path = window.location.pathname
      switch (path) {
        case (path.match(/\/play\/.*/) || {}).input:
          Child = <Play user={this.state && this.state.user ? this.state.user : null} game={this.state && this.state.game ? this.state.game : null} />
          break;
        case (path.match(/\/edit\/.*/) || {}).input:
          Child = <Edit />
          break;
        case '/decks':
          Child = <Decks />
          break;
        case (path.match(/\/create\/?.*/) || {}).input:
          Child = <CreateGame user={this.state && this.state.user ? this.state.user : null} />
          break;
        case '/init':
          Child = <InitializeGame />
          break;
        default:
          Child = <FindGame />
      }
    return Child;
  }

  render() {
    console.log(this.state)
    const game = this.state && this.state.game ? this.state.game : null;
    const user = this.state && this.state.user ? this.state.user : null;
    return (
      <div className="App">
        <Header game={game} user={user} />

        <div className="main">
        {this.getRoute()}
        </div>
      </div>
    );
  }
}

export default App;
