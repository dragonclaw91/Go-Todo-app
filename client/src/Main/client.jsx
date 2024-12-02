import React from "react";
import { Col, Row, InputGroup, Form, Button } from "react-bootstrap";
import "./client.css"
import Tasks from '../Tasks/tasks.jsx'
import { useStore, useInputStore,  postData, put } from "./client.controller.jsx"
import  { useState } from 'react';





function Main() {
    let tasks = []
    const { fetchData, postData, updateData, deleteData } = useStore();
    const data = useStore((state) => state.data);
    const isLoading = useStore((state) => state.isLoading);
    // const fetchData = useStore((state) => state.fetchData);
    //  const     inputValue = useInputStore((state) => state.inputValue);
    //   const   updateInputValue = useInputStore((state) => state.updateInputValue);

    
   const  [inputValue, updateInputValue] = useState('');
    const handleInputChange = (event) => {
        updateInputValue(event.target.value);
      };
    

    console.log("CHECKING DATA", data.data?.data.Value,inputValue)
    tasks = data.data?.data.Value || []



    // const handleClick = async () => {
    //     console.log('Button clicked!');
    //  const response = await post({"task":`${inputValue}`})
    //  console.log("RESPONSE 2",response)
    //     updateInputValue('');
    //   await  fetchData();
    //     // updateState(response.data); 
    // }

    const handleClick = async (id) => {
        try {
          await postData(inputValue); // POST the form data
          await fetchData();     
                updateInputValue('');   // Fetch updated data
        } catch (err) {
          console.error(err);
        }
      };

      const handlePut = async (id) => {
        try {
          await updateData(id); // POST the form data
          await fetchData();        // Fetch updated data
        } catch (err) {
          console.error(err);
        }
      };

      const handleDelete = async (id) => {
        console.log("DELETEING ID",id)
        try {
          await deleteData(id); // POST the form data
          await fetchData();        // Fetch updated data
        } catch (err) {
          console.error(err);
        }
      };

    React.useEffect(() => {
        fetchData();
    }, []);
    return (

        <Row >
            <Col>
                <Row className="center ">
                    Tasks
                </Row>
            </Col>
            <Col className="max" >
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
            <Col className="max">
                {console.log("TASKS", tasks)}
                {tasks.map((task) => {
                    return <Tasks key={task.id} task={task} delete={()=> {handleDelete(task.id)}} put={()=> {handlePut(task.id)}} />
                })}
            </Col>
        </Row>
    );
}


export default Main;