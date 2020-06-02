import React from 'react';
import logo from './logo.svg';
import './App.css';

//import the semantic-ui-react
import {Container} from "semantic-ui-react";

//import To DO list component'
import ToDoList from "./To-Do-List";

function App() {
  return (
    <div className="App">
      <Container>
        <ToDoList/>
      </Container>
    </div>
  );
}

export default App;
