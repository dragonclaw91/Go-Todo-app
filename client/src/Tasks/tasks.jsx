import React from "react";
import { Col, Row, InputGroup, Form, Button } from "react-bootstrap";
import "./tasks.css"
import { TbHttpDelete } from "react-icons/tb";
import { useState } from 'react';
import { useStore, put } from"../Main/client.controller.jsx"

function Tasks(props) {
    const [val, setVal] = useState(false);
    // const fetchData = useStore((state) => state.fetchData);
    // //  const     inputValue = useInputStore((state) => state.inputValue);
    // //   const   updateInputValue = useInputStore((state) => state.updateInputValue);
    // React.useEffect(() => {
    //     fetchData();
    // }, []);
    // const handleClick = () => {
    //     console.log('Button clicked!');
    //     put({"id":`${props.task.id}`})
    //     // fetchData();
    //     setVal(!val)
    // }

    return (

        <Row className="task space" >
            <Col>
                <TbHttpDelete size={30} color="red" />
            </Col>
            <Col className="test text" >
                {props.task.task}
            </Col>
            <Col  >
            {console.log("PROPS",props.task.iscompleted)}
                <label class="container">
                    <input type="checkbox" onChange={props.put}  />
                    <span class="checkmark"></span>
                </label>

            </Col>
        </Row>
    );
}


export default Tasks;