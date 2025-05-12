import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/(protected)/page')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/(protected)/page"!</div>
}
