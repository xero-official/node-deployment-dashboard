import React, { Component } from "react";
import { color } from './constants';

export default class TopNav extends Component {
    render() {
        return (<div style={{ height: '40px' }}>
            <h3 style={{
                textAlign: 'right',
                backgroundColor: color,
                padding: '5px',
                color: 'white',
                fontSize: '10pt'
            }}>Settings</h3>
        </div>)
    }
}