'use client';

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useHookCreateTodo } from "@/hooks/todos/useCreateTodo";
import { useHookGetTodos } from "@/hooks/todos/useGetTodos";
import { FormEvent } from "react";



export function HomePageClient() {
  const { mutate, isPending } = useHookCreateTodo();
  const { data, isLoading } = useHookGetTodos();
  const submitHandler = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const formData = new FormData(e.currentTarget)
    const title = formData.get("title") as string;
    const email = formData.get("email") as string;
    console.log({ email, title })
    mutate({ title, email })
  }

  return <div className="w-150 p-5">
    <form className="flex flex-col gap-2" onSubmit={(e) => submitHandler(e)}>
      <h2 className="text-2xl font-bold">Todo App</h2>
      <div>
        <Label>Title</Label>
        <Input name="title" title="title" type="text" />
      </div>
      <div>
        <Label>Email</Label>
        <Input name="email" title="email" type="text" />
      </div>
      <Button type="submit">{isPending ? "creating ..." : "create"}</Button>
    </form>
    {
      isLoading ?
        <div>loading...</div>
        :
        <pre>
          {data ? JSON.stringify(data, null, 2) : "todos empty"}
        </pre>
    }
  </div>
}
