import React, { Component } from 'react';
import GameActions from '../actions/game';
// import GameStore from '../stores/game';



class FinalScore extends Component {


	componentWillMount() {
	}

	componentDidMount() {
		GameActions.unsetGame()
	}

	renderFinalScore() {
		console.log(this.props.game.finalScore)
		return (<div>TODO</div>);

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