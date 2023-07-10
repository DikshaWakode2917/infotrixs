import React from "react";
import "./App.css";
import "./TodoList.css";
// import the Container Component from the semantic-ui-react
import { Container } from "semantic-ui-react";
// import the ToDoList component
import ToDoList from "./component/TodoList";
function App() {
  return (
    <div>
      <Container>
        <ToDoList></ToDoList>
      </Container>
    </div>
  );
}
export default App;