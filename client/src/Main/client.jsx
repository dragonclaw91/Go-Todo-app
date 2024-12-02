import React from "react";
import { Col, Row, InputGroup, Form, Button } from "react-bootstrap";
import "./client.css"
import Tasks from '../Tasks/tasks.jsx'
import { useStore, useInputStore, post, put } from "./client.controller.jsx"
import  { useState } from 'react';





function Main() {
    let tasks = []
    // const { fetchData, data } = useStore();
    const data = useStore((state) => state.data);
    const isLoading = useStore((state) => state.isLoading);
    const fetchData = useStore((state) => state.fetchData);
    //  const     inputValue = useInputStore((state) => state.inputValue);
    //   const   updateInputValue = useInputStore((state) => state.updateInputValue);

    
   const  [inputValue, updateInputValue] = useState('');
    const handleInputChange = (event) => {
        updateInputValue(event.target.value);
      };
    

    console.log("CHECKING DATA", data.data?.data.Value,inputValue)
    tasks = data.data?.data.Value || []
    const handleClick = async () => {
        console.log('Button clicked!');
     const response = await post({"task":`${inputValue}`})
     console.log("RESPONSE 2",response)
        updateInputValue('');
      await  fetchData();
        // updateState(response.data); 
    }

    const  handleput =  (id) => {
        console.log('PUT Button clicked!',id);
       put({"id":`${id}`})
      console.log("FINISHED PUT")
      fetchData();
      console.log("FINISHED FETCH")
    }
    React.useEffect(() => {
        fetchData();
    }, []);
    return (

        <Row>
            <Col>
                <Row className="center ">
                    Tasks
                </Row>
            </Col>
            <Col  >
                <InputGroup className="input" >
                    <Form.Control type="text"
                        value={inputValue}
                        onChange={(e) => handleInputChange(e) } aria-label="Text input with checkbox">
                    </Form.Control>
                    <Row className="btn_pad">
                        <Button onClick={(e) => handleClick(e)} className="button"  >
                            +
                        </Button>
                    </Row>
                </InputGroup>
            </Col>
            <Col>
                {console.log("TASKS", tasks)}
                {tasks.map((task) => {
                    return <Tasks key={task.id} task={task} put={()=> {handleput(task.id)}} />
                })}
            </Col>
        </Row>
    );
}


export default Main;