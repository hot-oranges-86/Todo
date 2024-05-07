import { useEffect, useState } from "react";
import "./App.css";
import { fetchTodos } from "./services/todos";
import { TodoList } from "./components/TodoList";
import { AddTodo } from "./components/AddTodo";
import { Routes, Route, Link } from "react-router-dom";
import { TodoPage } from "./pages/TodoPage";

function App() {
  const [todos, setTodos] = useState([]);

  useEffect(() => {
    fetchTodos().then(setTodos);
  }, []);

  // return (
  //   <>
  //     <nav>
  //       <ul>
  //         <li>
  //           <Link to="/todos">Home</Link>
  //         </li>
  //         <li>
  //           <Link to="/todos/add">Add To-Do</Link>
  //         </li>
  //       </ul>
  //     </nav>

  //     <Routes>
  //       <Route path="/todos">
  //         <Route index element={<TodoList todos={todos} />} />
  //         <Route path="add" element={<AddTodo />} />
  //         <Route path=":id" element={<TodoPage />} />
  //       </Route>
  //     </Routes>
  //   </>
  // );

  return (
    <>
      <AddTodo />
      <TodoList todos={todos} />
    </>
  );
}

export default App;
