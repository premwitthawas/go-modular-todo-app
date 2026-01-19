import { useQuery } from "@tanstack/react-query";
import { useHookTodoAPI } from "./usetodoApi";

export function useHookGetTodos() {
  const { getTodos } = useHookTodoAPI()
  return useQuery({
    queryKey: ['todos'],
    queryFn: getTodos,
    staleTime: 1000 * 60 * 5,
    retry: 2
  })
}
