import React, { Component } from "react";
import { events, color, options, nodeTypes } from "./constants";
import {
    Container, Row, Col, Button, ListGroup, InputGroup, FormControl
} from 'react-bootstrap'

export default class Body extends Component {

    constructor(props) {
        super(props);
        this.state = {
            contractOption: null,
            log: 'System logs',
            nodeType: null,
            text: ''
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
            if (obj.event === 'LOG')
                this.setState({ log: this.state.log + '. ' + obj.log });
        }
    }

    onNodeTypeSelected(nodeType) {
        this.setState({ nodeType });
        if (this.state.contractOption === options.ADD_NODE) {
            this.ws.send(JSON.stringify({
                event: events.SELECT_OPTION,
                contractOption: this.state.contractOption,
                nodeType
            }))
            this.setState({ contractOption: null, nodeType: null, text: '' });
        }
    }

    onProceedSelected() {
        const { contractOption, nodeType, text } = this.state;
        if (text === '') {
            alert('Please fill');
            return;
        }
        this.ws.send(JSON.stringify({
            event: events.SELECT_OPTION,
            contractOption, nodeType, privateKey: text, nodeAddress: text
        }))
        this.setState({ contractOption: null, nodeType: null, text: '' });
    }

    onOptionSelected(contractOption) {
        this.setState({ contractOption, nodeType: null, text: '' });
    }

    render() {
        const { contractOption, nodeType, log, text } = this.state;

        const secondAreaVisible = contractOption === options.ADD_NODE || contractOption === options.LOOKUP_NODE || contractOption === options.REMOVE_NODE;
        const thirdAreaVisible = (contractOption === options.LOOKUP_NODE || contractOption === options.REMOVE_NODE) && nodeType;
        const size = thirdAreaVisible ? [4, 4, 4] : secondAreaVisible ? [4, 6, 0] : [6, 0, 0];

        const label = contractOption === options.LOOKUP_NODE ? 'Lookup Address' : 'Private key';
        return (<Container style={{ maxWidth: '100%', marginTop: '0px' }}>
            <Row>
                <Col xs={size[0]}>
                    <div style={{ marginTop: '10px', height: '380px' }}>
                        <div style={{ border: '1', borderRadius: '10px', height: '380px', backgroundColor: '#ededed' }}>
                            <Button onClick={e => this.onOptionSelected(options.ADD_NODE)} style={{ backgroundColor: color, width: 'calc(100% - 20px)', margin: '10px' }} variant="primary">Add a Node</Button>
                            <Button onClick={e => this.onOptionSelected(options.REMOVE_NODE)} style={{ backgroundColor: color, width: 'calc(100% - 20px)', margin: '10px' }} variant="primary">Remove a Existing node</Button>
                            <Button onClick={e => this.onOptionSelected(options.LOOKUP_NODE)} style={{ backgroundColor: color, width: 'calc(100% - 20px)', margin: '10px' }} variant="primary">Lookup Node</Button>
                            <Button onClick={e => { }} style={{ backgroundColor: color, width: 'calc(100% - 20px)', margin: '10px' }} variant="primary">Exit</Button>
                        </div>
                    </div>
                </Col>
                {secondAreaVisible ?
                    <Col xs={size[1]}>
                        <div style={{ marginTop: '10px', height: '380px' }}>
                            <div style={{ marginTop: '10px', border: '1', borderRadius: '10px', height: '380px', backgroundColor: '#ededed' }}>
                                <div style={{ color: 'white', width: 'calc(100% - 20px)', margin: '10px' }} >.</div>
                                <div style={{ width: 'calc(100% - 20px)', margin: '10px' }} >Select Node Type</div>
                                <ListGroup style={{ marginLeft: '10px', marginRight: '10px' }} >
                                    {nodeTypes.map((nodeType, i) => <ListGroup.Item action onClick={() => this.onNodeTypeSelected(i + 1)}>
                                        {(i + 1) + '. ' + nodeType}</ListGroup.Item>)}
                                    <ListGroup.Item action onClick={e => this.setState({ nodeType: null, contractOption: null })}>Cancel</ListGroup.Item>
                                </ListGroup>
                            </div>
                        </div>
                    </Col> : null}
                {thirdAreaVisible ?
                    <Col xs={size[2]}>
                        <div style={{ marginTop: '10px', height: '380px' }}>
                            <div style={{ marginTop: '10px', border: '1', borderRadius: '10px', height: '380px', backgroundColor: '#ededed' }}>
                                <div style={{ color: 'white', width: 'calc(100% - 20px)', margin: '10px' }} >.</div>
                                <div style={{ width: 'calc(100% - 20px)', margin: '10px' }} >Please type {label}</div>
                                <InputGroup style={{ width: 'calc(100% - 20px)', paddingTop: '10px', margin: '10px' }} className="mb-3">
                                    <InputGroup.Prepend>
                                        <InputGroup.Text id="basic-addon1">{label}</InputGroup.Text>
                                    </InputGroup.Prepend>
                                    <FormControl
                                        onChange={e => this.setState({ text: e.target.value })}
                                        value={text}
                                        placeholder={"Type " + label}
                                        aria-label="data-thing"
                                        aria-describedby="basic-addon1"
                                    />
                                </InputGroup>
                                <Button onClick={e => this.onProceedSelected()} style={{ backgroundColor: color, width: 'calc(100% - 20px)', margin: '10px' }} variant="primary">Proceed</Button>
                                <Button onClick={e => this.setState({ nodeType: null, contractOption: null })} style={{ borderColor: color, color: color, width: 'calc(100% - 20px)', margin: '10px' }} variant="outline-primary">Cancel</Button>

                            </div>
                        </div>
                    </Col> : null}
            </Row>
            <Row>
                <Col xs={12} style={{ marginTop: '10px' }}>
                    <div style={{ marginTop: '10px', height: '350px' }}>
                        <div style={{ paddingLeft: '10px', overflow: 'hidden', border: '1', borderRadius: '10px', height: '380px', backgroundColor: '#ededed' }}>
                            <div style={{ color: 'white', width: 'calc(100% - 20px)', margin: '0px' }} >.</div>
                            {log}
                        </div>
                    </div>
                </Col>
            </Row>
        </Container >)
    }
}