import axios from 'axios';
import React, { Component } from 'react';
import './App.css';
import SearchBar from './SearchBar';
import SpyList from './SpyList';
import Loading from './Loading';

class App extends Component {
  constructor() {
    super();
    this.state = {
      'error': null,
      'loading': true,
      'spies': [],
      'spiesForSearch': []
    };
    this.raw = [];
  }

  componentDidMount() {
    axios.get('/spies')
    .then((r) => {
      this.raw = r.data;
      this.setState({
        loading: false,
        spies: r.data,
        spiesForSearch: r.data
      })
    })
    .catch((err) => {
      let msg = '';
      try {
        let content = JSON.parse(err.response.data);
        msg = content['message'];
      } catch (e) {
        msg = err.response.data;
      }
      this.setState({
        error: {
          message: msg,
          responseCode: err.response.status
        }
      });
    });
  }

  onSearch(results) {
    this.setState({
      spies: results
    });
  }

  onSearchReset() {
    this.setState({
      spies: this.raw
    });
  }

  render() {
    let content;
    if (this.state.loading === true) {
      content = (
        <Loading/>
      );
    } else {
      content = (
        <SpyList spies={this.state.spies}/>
      );
    }

    if (this.state.error !== null) {
      content = (
        <div className="alert alert-danger" role="alert">
          HTTP Response Code: {this.state.error.responseCode}<br/>
          Message: {this.state.error.message}
        </div>
      );
    }

    return (
      <div className="App">
        <header className="mb-4">
          <nav className="navbar navbar-expand-lg navbar-light bg-light">
            <span className="navbar-brand">Image Spy</span>
            <SearchBar spies={this.state.spiesForSearch} onSearch={this.onSearch.bind(this)} onReset={this.onSearchReset.bind(this)} />
            <a className="nav-item nav-link" href="https://github.com/imagespy/imagespy">GitHub</a>
          </nav>
        </header>
        <div className="container-fluid">
          {content}
        </div>
      </div>
    );
  }
}

export default App;
