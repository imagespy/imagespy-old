import React, { Component } from 'react';
import Spy from './Spy';

import './SpyList.css';

class SpyList extends Component {
  renderSpies() {
    return this.props.spies.sort().map((spy) => {
      return (
        <Spy key={spy.labels.container} spy={spy}/>
      )
    })
  }

  render() {
    let spies = this.renderSpies();
    return (
      <div className="SpyList">
        <div className="row mb-3">
          <div className="col-12 mx-auto">
            <table className="table">
              <thead>
                <tr>
                  <th>Container</th>
                  <th>Image</th>
                  <th>Current Tag</th>
                  <th>Latest Tag</th>
                </tr>
              </thead>
              <tbody>
                {spies}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    );
  }
}

export default SpyList;
