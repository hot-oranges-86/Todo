import { DeleteTodo } from "./DeleteTodo";
import { ToggleTodo } from "./ToggleTodo";
import "../styles/list.css";

export const TodoItem = ({ todo, setTodos }) => {
  return (
    <li key={todo.id}>
      <ToggleTodo {...todo} setTodos={setTodos} />

      <h3>{todo.name}</h3>

      <DeleteTodo id={todo.id} setTodos={setTodos} />
    </li>
  );
};
