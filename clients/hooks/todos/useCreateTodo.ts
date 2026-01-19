import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useHookTodoAPI } from "./usetodoApi";
import type { CreateTodoReq, TodoRes } from "@/types/todos";



export function useHookCreateTodo() {
  const queryclient = useQueryClient();
  const useHookAPI = useHookTodoAPI();
  return useMutation({
    mutationFn: (req: CreateTodoReq) => useHookAPI.createTodo(req),
    onSuccess: (data: TodoRes) => {
      queryclient.invalidateQueries({ queryKey: ['todos'] })
      console.log("Created successfully:", data);
    },
    onError: (error) => {
      console.error("Mutation Error:", error.message);
    }
  })
};
