export const API_TODOS_URL = "http://localhost:8080/todos";

export const fetchTodos = () => fetch(API_TODOS_URL).then((res) => res.json());

// export const fetchTodos = () => async (dispatch) => {
//   const response = await fetch(API_TODOS_URL);
// };
