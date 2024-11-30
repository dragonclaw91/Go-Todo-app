import React from "react";
import { Col, Row, InputGroup, Form, Button } from "react-bootstrap";
import "./tasks.css"
import { TbHttpDelete } from "react-icons/tb";



function Tasks() {
    return (

        <Row className="task space" >
            <Col>
            <TbHttpDelete size={30} color="red" />
            </Col>
            <Col className="test text" >
                Create New Project
            </Col>
            <Col  >
                <label class="container">
                    <input type="checkbox" checked="checked" />
                    <span class="checkmark"></span>
                </label>

            </Col>
        </Row>
    );
}


export default Tasks;