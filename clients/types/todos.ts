export interface CreateTodoReq {
  title: string
  email: string
}

export interface TodoRes {
  id: string
  title: string
  email: string
  done: boolean
  created_at: Date
  updated_at: Date
}
