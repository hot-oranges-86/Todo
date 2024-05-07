import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";

export const TodoPage = () => {
  const { id } = useParams();

  const [todo, setTodo] = useState(null);

  return <h1>{1}</h1>;
};
