import React, { Component } from 'react';
import GameActions from '../actions/game';

class FinalScore extends Component {


	componentDidMount() {
		GameActions.unsetGame()
	}

	renderFinalScore() {
		let players = [];

		for (const i in this.props.game.players) {
			if (!this.props.game.players[i]) {
				continue;
			}
			for (const id in this.props.game.finalScore) {
				if (!this.props.game.finalScore[id]) {
					continue;
				}
				if (id === this.props.game.players[i]._id) {
					players.push(<div key={'wins'+id}className="wins">Player: {this.props.game.players[i].name} ... Votes: {this.props.game.finalScore[id].length}</div> );
				}
			}
		}
		return (<div>{players}</div>);

	}

	render() {
		return (
			<div className="play">Final Score: 
				{this.renderFinalScore()}
			</div>
		);
	}
}

export default FinalScore;