import axios from 'axios';
import React, { Component } from 'react';
import './App.css';
import SearchBar from './SearchBar';
import SpyList from './SpyList';

class App extends Component {
  constructor() {
    super();
    this.state = {
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
        spies: r.data,
        spiesForSearch: r.data
      })
    })
    .catch((err) => {
      console.log(err)
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
          <SpyList spies={this.state.spies}/>
        </div>
      </div>
    );
  }
}

export default App;
