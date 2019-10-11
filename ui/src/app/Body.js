import React, { Component } from "react";
import { events, color, options, nodeTypes } from "./constants";
import {
    Container, Row, Col, Button, ListGroup, Text
} from 'react-bootstrap'

export default class Body extends Component {

    constructor(props) {
        super(props);
        this.state = {
            selectedOption: null
        }

        try {
            this.ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");
        } catch (error) {
            console.log(error);
        }
    }

    componentDidMount() {

        this.ws.onmessage = (message) => {
            const obj = JSON.parse(message.data);
            console.log(obj.event);
            if (obj.event === 'TEST')
                this.setState({});
            console.log(obj.payload);
        }
    }

    onNodeTypeSelected(nodeType) {
        this.ws.send(JSON.stringify({
            event: events.SELECT_OPTION,
            contractOption: this.state.selectedOption,
            nodeType
        }))
        this.setState({ nodeType });
    }

    onOptionSelected(selectedOption) {
        this.setState({ selectedOption });
    }

    render() {
        const { selectedOption } = this.state;

        return (<Container style={{ maxWidth: '100%', marginTop: '0px' }}>
            <Row>
                <Col xs={4}>
                    <div style={{ marginTop: '10px', height: '350px' }}>
                        <div style={{ border: '1', borderRadius: '10px', height: '380px', backgroundColor: '#ededed' }}>
                            <Button onClick={e => this.onOptionSelected(options.ADD_NODE)} style={{ backgroundColor: color, width: 'calc(100% - 20px)', margin: '10px' }} variant="primary">Add a Node</Button>
                            <Button onClick={e => this.onOptionSelected(options.REMOVE_NODE)} style={{ backgroundColor: color, width: 'calc(100% - 20px)', margin: '10px' }} variant="primary">Remove a Existing node</Button>
                            <Button onClick={e => this.onOptionSelected(options.LOOKUP_NODE)} style={{ backgroundColor: color, width: 'calc(100% - 20px)', margin: '10px' }} variant="primary">Lookup Node</Button>
                            <Button onClick={e => this.onOptionSelected(options.EXIT)} style={{ backgroundColor: color, width: 'calc(100% - 20px)', margin: '10px' }} variant="primary">Exit</Button>
                        </div>
                    </div>
                </Col>
                {selectedOption === options.ADD_NODE ?
                    <Col xs={8}>
                        <div style={{ marginTop: '10px', height: '400px' }}>
                            <div style={{ marginTop: '10px', border: '1', borderRadius: '10px', height: '380px', backgroundColor: '#ededed' }}>
                                <div style={{ width: 'calc(100% - 20px)', margin: '10px' }} >Select Node Type</div>
                                <ListGroup style={{ marginLeft: '10px', marginRight: '10px' }} >
                                    {nodeTypes.map((nodeType, i) => <ListGroup.Item action onClick={() => this.onNodeTypeSelected(i + 1)}>
                                        {(i + 1) + '. ' + nodeType}</ListGroup.Item>)}
                                </ListGroup>
                            </div>
                        </div>
                    </Col> : null}
            </Row>
        </Container >)
    }
}