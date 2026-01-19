'use client'

import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import { ReactQueryDevtools } from "@tanstack/react-query-devtools"
import { ReactNode } from "react"

interface Props {
  children: ReactNode
}

export function TanstackQueryProvider({ children }: Props) {
  const queryClient = new QueryClient()
  return <QueryClientProvider client={queryClient}>
    {children}
    <ReactQueryDevtools initialIsOpen />
  </QueryClientProvider>
}
