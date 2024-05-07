import { API_TODOS_URL, fetchTodos } from "../services/todos";

export const DeleteTodo = ({ id, setTodos }) => {
  const handleClick = async () => {
    try {
      await fetch(`${API_TODOS_URL}/${id}`, {
        method: "DELETE",
      });

      const updatedTodos = await fetchTodos();
      setTodos(updatedTodos);
    } catch {
      console.error("Delte error", error);
    }
  };

  return <button onClick={handleClick}>DELETE</button>;
};
