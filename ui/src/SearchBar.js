import Fuse from 'fuse.js';
import React, { Component } from 'react';

import './SearchBar.css';

class SearchBar extends Component {
  constructor() {
    super();
    this.fuse = null;
  }

  componentDidMount() {
    var options = {
      threshold: 0.1,
      location: 0,
      distance: 100,
      maxPatternLength: 32,
      minMatchCharLength: 2,
      keys: [
        "labels.container",
        "name"
      ]
    };
    this.fuse = new Fuse([], options);
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps !== this.props) {
      this.fuse.setCollection(nextProps.spies);
    }
  }

  handleChange(event) {
    let value = event.target.value;
    if (value === '') {
      this.props.onReset();
    } else {
      let results = this.fuse.search(event.target.value);
      this.props.onSearch(results);
    }
  }

  render() {
    return (
      <form className="mx-2 my-auto d-inline w-100">
        <div className="input-group">
          <input onChange={this.handleChange.bind(this)} className="form-control" type="text" placeholder="Search for container or image" aria-label="Search" />
        </div>
      </form>
    )
  }
}

export default SearchBar;
