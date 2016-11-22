import React, { Component } from 'react';
import '../css/findGame.css'

class Directions extends Component {

	render() {
		return (
			<div className="directions">
				<p>If your 'friend' gave you a big, long ID number, paste it here to find the game everyone's playing.
				If you're in charge, click 'start a game' above. Then, email or text or whatever, the ID that comes up.
				</p>
				<h5>The Difference Between works in two rounds:</h5>
				<ol>
					<li>Each player reads the "difference between" phrase with the two random dealer cards. Said players play the card from their own 
					hand that makes things funniest...or truest.</li>
					<li>Each player then votes for the funniest entry that round. Round scores are displayed after everyone votes.</li>
				</ol>
			</div>
		);
	}
}

export default Directions;