import { TodoItem } from "./TodoItem";

export const TodoList = ({ todos = [] }) => {
  return (
    <div>
      <ul>
        {todos.map((todo) => (
          <li key={todo.id}>
            <TodoItem name={todo.name} />
          </li>
        ))}
      </ul>
    </div>
  );
};
