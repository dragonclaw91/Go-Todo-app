import React from "react";
import { Col, Row, InputGroup, Form, Button } from "react-bootstrap";
import "./client.css"
import Tasks from '../Tasks/tasks.jsx'


function Main() {
    return (

        <Row>
            <Col>
                <Row className="center ">
                    Tasks
                </Row>
            </Col>
            <Col  >
                <InputGroup className="input" >
                    <Form.Control aria-label="Text input with checkbox">
                    </Form.Control>
                    <Row className="btn_pad">
                        <Button className="button"  >
                            +
                        </Button>
                    </Row>
                </InputGroup>
            </Col>
            <Col>
            <Tasks>
            </Tasks>
            </Col>
        </Row>
    );
}


export default Main;