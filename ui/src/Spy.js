import classNames from 'classnames';
import moment from 'moment';
import React, { Component } from 'react';

import './Spy.css';

class Spy extends Component {
  chooseStatusClasses(spy) {
    if (spy.current_image.tag === spy.latest_image.tag) {
      if (moment(spy.created).isBefore(spy.latest_image.created)) {
        return ['table-danger'];
      }

      return ['table-success'];
    } else {
      return ['table-danger'];
    }
  }

  formatCreated(created) {
    return moment(created).utc().format('YYYY-MM-DD');
  }

  formatImageName(name) {
    let result = name;
    if (name.startsWith('library/')) {
      result = name.split('/')[1];
    }

    if (result.indexOf(':') !== -1) {
      result = result.split(':')[0]
    }

    return result;
  }

  render() {
    let spy = this.props.spy;
    let classes = classNames('Spy', this.chooseStatusClasses(spy));
    return (
      <tr className={classes}>
        <td>{spy.labels.container}</td>
        <td>{this.formatImageName(spy.name)}</td>
        <td>{spy.current_image.tag} <span>({this.formatCreated(spy.current_image.created)})</span></td>
        <td>{spy.latest_image.tag} <span>({this.formatCreated(spy.latest_image.created)})</span></td>
      </tr>
    )
  }
}

export default Spy;
