import { CreateTodoReq, TodoRes } from "@/types/todos";

const apiUri = process.env.NEXT_PUBLIC_API_URL as string;

async function createTodo(req: CreateTodoReq): Promise<TodoRes> {
  const res = await fetch(`${apiUri}/api/v1/todos`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(req),
    credentials: "include",
  });

  if (!res.ok) {
    const errorData = await res.json().catch(() => ({}));
    throw new Error(errorData.message || `Error: ${res.status}`);
  }

  return await res.json();
}

async function getTodos(): Promise<TodoRes[]> {
  const res = await fetch(`${apiUri}/api/v1/todos`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });

  if (!res.ok) {
    const errorData = await res.json().catch(() => ({}));
    throw new Error(errorData.message || `Error: ${res.status}`);
  }

  return await res.json();
}

export function useHookTodoAPI() {
  return { createTodo, getTodos }
}
