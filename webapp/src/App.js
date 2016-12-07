import React, { Component } from 'react';
import './App.css';
import Play from './components/play';
import Decks from './components/decks';
import Edit from './components/edit';
import CreateGame from './components/createGame';
import InitializeGame from './components/initializeGame';
import FindGame from './components/findGame';
import Header from './components/header';
import EditPlayer from './components/editPlayer';
import EditCard from './components/editCard';
import Audit from './components/audit';
import Error from './components/error';

class App extends Component {

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
    const path = window.location.pathname;
      switch (path) {
        case (path.match(/\/play\/.*/) || {}).input:
          Child = <Play  />;
          break;
        case (path.match(/\/edit\/.*/) || {}).input:
          Child = <Edit />;
          break;
        case '/decks':
          Child = <Decks />;
          break;
        case (path.match(/\/create\/?.*/) || {}).input:
          Child = <CreateGame />;
          break;
        case '/init':
          Child = <InitializeGame />;
          break;
        case '/player':
          Child = <EditPlayer />;
          break;
        case '/card':
          Child = <EditCard />;
          break;
        case '/audit':
          Child = <Audit />;
          break;
        default:
          Child = <FindGame />;
      }
    return Child;
  }

  render() {
    return (
      <div className="App">
        <Header  />

        <div className="main">
        {this.getRoute()}
        </div>
        <Error />
      </div>
    );
  }
}

export default App;
