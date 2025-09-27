import { Button } from "@/components/ui/button"
import React, { useEffect } from "react"
import { Badge } from "@/components/ui/badge"
import {
  Card,
  // CardAction,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import ProjectTile from "./components/ProjectTile"
import NewProjectTile from "./components/NewProjectTile"

function Home() {
  const [count, setCount] = React.useState(0)

  return (
    <div className="min-h-screen bg-white py-8 px-4">
      <div className="flex flex-col">
        <div className="mb-6">
          <h1 className="text-3xl font-bold">Projekty</h1>
          <p className="text-muted-foreground mt-2">
            Zarządzaj swoimi projektami i segreguj treści
          </p>
        </div>
        <div className="flex flex-col gap-4">
          <ProjectTile year="2024" title="USA West" />
          <ProjectTile year="2025" title="Sri Lanka" />
          <NewProjectTile />
        </div>
      </div>
    </div>
  )
}

export default Home
