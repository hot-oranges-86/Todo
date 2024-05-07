import { useState } from "react";
import { API_TODOS_URL, fetchTodos } from "../services/todos";

export const AddTodo = ({ setTodos }) => {
  const [name, setName] = useState("");
  const [completed, setCompleted] = useState(false);

  const handleSubmit = async (ev) => {
    ev.preventDefault();

    const response = await fetch(API_TODOS_URL, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: name,
        completed: completed,
      }),
    });

    setName("");

    const updatedTodos = await fetchTodos();
    setTodos(updatedTodos);
  };

  return (
    <div>
      <h2>Create ToDo Item</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label>
            Name:
            <input
              type="text"
              value={name}
              onChange={(ev) => setName(ev.target.value)}
              required
            />
          </label>
        </div>
        <div>
          <label>
            Completed:
            <input
              type="checkbox"
              checked={completed}
              onChange={(ev) => setCompleted(ev.target.checked)}
            />
          </label>
        </div>
        <button type="submit">Create</button>
      </form>
    </div>
  );
};
