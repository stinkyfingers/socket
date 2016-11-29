import React, { Component } from 'react';
import ChatStore from '../stores/chat';
import ChatActions from '../actions/chat';
import '../css/chat.css'

class Chat extends Component {
  constructor() {
    super()
    this.statusUpdate = this.statusUpdate.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  statusUpdate(status) {
    if (status.messages) {
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

  handleSubmit(e) {
    if (e.charCode !== 13) {
      return;
    }
    ChatActions.send(this.props.game._id, this.state.message, this.props.user);
    this.refs.input.value = '';
  }

  renderMessages() {
    const maxLen = 20;
    const messages = [];
    if (this.state.messages.length > maxLen) {
      this.state.messages.splice(0, this.state.messages.length - maxLen);
    }
    for (let i in this.state.messages) {
      if (!this.state.messages[i]) {
        continue;
      }
      messages.push(<div className="chatLine" key={'message' + i}><span className="bold">{this.state.messages[i].message.playerName}:</span> {this.state.messages[i].message.text}</div>)
    }

    return (
      <div className="chatWindow">
        {messages}
      </div>
    );
  }

  render() {

    return (
      <div className="chat">
        <input type="text" name="message" onKeyPress={this.handleSubmit} onChange={this.handleChange} ref="input" placeholder="chat..." />
        <button onClick={this.handleSubmit} className="btn submit">Send</button>
        {this.state && this.state.messages ? this.renderMessages() : null}
      </div>
    );
  }
}

export default Chat;
