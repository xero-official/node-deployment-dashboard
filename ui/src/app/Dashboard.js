import React, { Component } from "react";
import TopNav from "./TopNav";
import Body from "./Body";

export default class App extends Component {
    render() {
        return <div>
            <TopNav />
            <Body />
        </div>
    }
}