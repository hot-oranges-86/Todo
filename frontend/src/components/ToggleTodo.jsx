import { API_TODOS_URL, fetchTodos } from "../services/todos";

export const ToggleTodo = ({ id, name, completed, setTodos }) => {
  const handleChange = async () => {
    const updatedCompleted = !completed;
    await fetch(`${API_TODOS_URL}/${id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id: id,
        name: name,
        completed: updatedCompleted,
      }),
    });

    const updatedTodos = await fetchTodos();
    setTodos(updatedTodos);
  };

  return <input type="checkbox" checked={completed} onChange={handleChange} />;
};
