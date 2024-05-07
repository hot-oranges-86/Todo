export const API_TODOS_URL = "http://localhost:8080/todos";

export const fetchTodos = () => fetch(API_TODOS_URL).then((res) => res.json());
