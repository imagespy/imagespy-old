import React, { Component } from 'react';

import './Loading.css';

class Loading extends Component {
  render() {
    return (
      <div className="Loading row mb-3 justify-content-center">
        <div className="col-12 mx-auto">
          <p>
            Loading image spies for the first time...
          </p>
          <p>
            This can take a few seconds.
          </p>
        </div>
      </div>
    )
  }
}

export default Loading
