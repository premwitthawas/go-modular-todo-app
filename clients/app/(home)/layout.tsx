'use client'

import { TanstackQueryProvider } from "@/providers/tanstack-query-provider"
import { ReactNode } from "react"

interface HomelayoutProps {
  children: ReactNode
}

export default function HomeLayout({ children }: HomelayoutProps) {
  return <>
    <TanstackQueryProvider>
      {children}
    </TanstackQueryProvider>
  </>
}
