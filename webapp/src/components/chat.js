import React, { Component } from 'react';
import ChatStore from '../stores/chat';
import ChatActions from '../actions/chat';
import '../App.css'

class Chat extends Component {
  constructor() {
    super()
    this.statusUpdate = this.statusUpdate.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  statusUpdate(status) {
    if (status.messages) {
      console.log(status)
      this.setState({ messages: status.messages });
    }
  }

  componentWillMount() {
    ChatActions.connect(this.props.game._id);
  }

  componentDidMount() {
    this.unlisten = ChatStore.listen(this.statusUpdate);
  }

  componentWillUnmount() {
    this.unlisten();
  }

  handleChange(e) {
    this.setState({ message: e.target.value });
  }

  handleSubmit() {
    ChatActions.send(this.props.game._id, this.state.message, this.props.user);
  }

  render() {
    return (
      <div className="chat">Chat

        <input type="text" name="message"  onChange={this.handleChange} />
        <button onClick={this.handleSubmit} className="btn submit">Send</button>
      </div>
    );
  }
}

export default Chat;
