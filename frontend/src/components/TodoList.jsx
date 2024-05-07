import { TodoItem } from "./TodoItem";

export const TodoList = ({ todos = [], setTodos }) => {
  return (
    <ul>
      {todos.map((todo) => (
        <TodoItem todo={todo} setTodos={setTodos} />
      ))}
    </ul>
  );
};
